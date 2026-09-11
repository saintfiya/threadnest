package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

//本项目使用简化版的投票分数
//投一票就加432分 86400/200 --> 200张赞成票可以给帖子续一天

/* 投票的几种情况
direction=1时,有两种情况:
	1.之前没有投过票,现在投赞成票 -->更新分数和投票记录 差值的绝对值:1 +432
	2.之前投反对票,现在改投赞成票 -->更新分数和投票记录 差值的绝对值:2 +432
direction=0时,有两种情况:
	1.之前投过赞成票,现在要取消投票 -->更新分数和投票记录 差值的绝对值:1 -432
	2.之前投过反对票,现在要取消投票 -->更新分数和投票记录 差值的绝对值:1 +432
direction=-1时,有两种情况:
	1.之前没有投过票,现在投反对票 -->更新分数和投票记录 差值的绝对值:1 -432
	2.之前投赞成票,现在改投反对票 -->更新分数和投票记录 差值的绝对值:2 -432*2

投票的限制（下述回写 MySQL 的归档流程尚未实现；个人投票 key 已按剩余窗口设置 TTL）:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期就不允许再投票了.
	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除那个 KeyPostVotedZsetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrPostNotExist   = errors.New("帖子不存在")
	ErrInvalidVote    = errors.New("投票值必须是 -1、0 或 1")
)

var voteScript = redis.NewScript(`
local post_time = redis.call('ZSCORE', KEYS[1], ARGV[1])
if not post_time then
  return -1
end

local now = tonumber(ARGV[4])
local max_age = tonumber(ARGV[5])
local age = now - tonumber(post_time)
if age > max_age then
  return -2
end

local old_vote = redis.call('ZSCORE', KEYS[3], ARGV[2])
if not old_vote then
  old_vote = 0
else
  old_vote = tonumber(old_vote)
end

local new_vote = tonumber(ARGV[3])
local delta = (new_vote - old_vote) * tonumber(ARGV[6])
redis.call('ZINCRBY', KEYS[2], delta, ARGV[1])

if new_vote == 0 then
  redis.call('ZREM', KEYS[3], ARGV[2])
else
  redis.call('ZADD', KEYS[3], new_vote, ARGV[2])
  local ttl = math.floor(max_age - age)
  if ttl > 0 then
    redis.call('EXPIRE', KEYS[3], ttl)
  end
end

return 0
`)

func CreatePost(ctx context.Context, postID int64) error {

	pipeline := client.TxPipeline()

	// 帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	_, err := pipeline.Exec(ctx)
	return err
}

func VoteForPost(ctx context.Context, userID, postID string, value float64) error {
	if value != -1 && value != 0 && value != 1 {
		return ErrInvalidVote
	}

	result, err := voteScript.Run(ctx, client, []string{
		getRedisKey(KeyPostTimeZset),
		getRedisKey(KeyPostScoreZset),
		getRedisKey(KeyPostVotedZsetPF + postID),
	}, postID, userID, value, time.Now().Unix(), oneWeekInSeconds, scorePerVote).Int()
	if err != nil {
		return err
	}
	switch result {
	case -1:
		return ErrPostNotExist
	case -2:
		return ErrVoteTimeExpire
	default:
		return nil
	}
}
