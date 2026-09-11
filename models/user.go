package models

import "time"

type User struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	UserID     int64     `gorm:"uniqueIndex;not null"`
	Username   string    `gorm:"type:varchar(64);uniqueIndex;not null"`
	Password   string    `gorm:"type:varchar(64);not null" json:"-"`
	Email      string    `gorm:"type:varchar(64)"`
	Gender     int8      `gorm:"not null;default:0"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	Token      string    `gorm:"-" json:"-"`
}
