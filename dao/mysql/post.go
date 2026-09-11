package mysql

import (
	"errors"
	"goweb/models"

	"gorm.io/gorm"
)

func CreatePost(p *models.Post) error {
	err := db.
		Create(p).
		Error

	if err != nil {

		return err
	}
	return nil
}

func DeletePost(postID int64) error {
	return db.Where("post_id = ?", postID).Delete(&models.Post{}).Error
}

// GetPostList 查询帖子列表
func PostList(page, size int64) (posts []*models.Post, err error) {
	err = db.
		Select("post_id, title, content, author_id, community_id, create_time").
		Order("create_time DESC").
		Offset(int((page - 1) * size)).
		Limit(int(size)).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	err = db.
		Where("post_id = ?", pid).
		First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrorPostNotExist
	}
	if err != nil {
		return nil, err
	}
	return
}
