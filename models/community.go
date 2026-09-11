package models

import "time"

type Community struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"-"`
	CommunityID   int       `gorm:"uniqueIndex;not null" json:"community_id"`
	CommunityName string    `gorm:"type:varchar(128);uniqueIndex;not null" json:"community_name"`
	Introduction  string    `gorm:"type:varchar(256);not null" json:"introduction"`
	CreateTime    time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime    time.Time `gorm:"autoUpdateTime" json:"-"`
}

type CommunityList struct {
	CommunityID   int    `json:"community_id"`
	CommunityName string `json:"community_name"`
}
