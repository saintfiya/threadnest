package logic

import "errors"

var (
	ErrPostNotExist      = errors.New("帖子不存在")
	ErrCommunityNotExist = errors.New("社区不存在")
	ErrVoteTimeExpire    = errors.New("投票时间已过")
	ErrInvalidVote       = errors.New("投票值必须是 -1、0 或 1")
)
