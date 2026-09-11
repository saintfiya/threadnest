package logic

import (
	"context"
	"errors"
	"goweb/dao/mysql"
	"goweb/dao/redis"
	"goweb/models"
	"goweb/pkg/snowflake"
)

func CreatePost(ctx context.Context, p *models.Post) error {
	//生成PostID
	p.PostID = snowflake.GenID()

	//保存到数据库
	err := mysql.CreatePost(p)
	if err != nil {
		return err
	}

	if err = redis.CreatePost(ctx, p.PostID); err != nil {
		// MySQL 已成功、Redis 失败时进行补偿，避免留下无法投票的帖子。
		if rollbackErr := mysql.DeletePost(p.PostID); rollbackErr != nil {
			return errors.Join(err, rollbackErr)
		}
		return err
	}

	return nil
}

func PostList(page int64, size int64) (posts []*models.Post, err error) {

	posts, err = mysql.PostList(page, size)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByID(pid int64) (post *models.Post, err error) {

	post, err = mysql.GetPostByID(pid)
	if err != nil {
		if errors.Is(err, mysql.ErrorPostNotExist) {
			return nil, ErrPostNotExist
		}
		return nil, err
	}
	return
}
