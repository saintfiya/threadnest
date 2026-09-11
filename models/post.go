package models

import "time"

type Post struct {
	Title       string    `gorm:"type:varchar(128);not null" binding:"required,max=128" json:"title"`
	Content     string    `gorm:"type:varchar(8192);not null" binding:"required,max=8192" json:"content"`
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"-"`
	PostID      int64     `gorm:"uniqueIndex;not null" json:"post_id"`
	AuthorID    int64     `gorm:"index;not null" json:"author_id"`
	CommunityID int64     `gorm:"index;not null" binding:"required,gt=0" json:"community_id"`
	Status      int8      `gorm:"not null;default:1" json:"status"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"-"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string             `json:"author_name"` // 作者
	VoteNum    int64              `json:"vote_num"`    // 投票数
	*Post                         // 嵌入帖子结构体
	*Community `json:"community"` // 嵌入社区信息
}
