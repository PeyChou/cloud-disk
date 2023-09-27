package models

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm.Model
	Identity           string `gorm:"type:varchar(36);"`
	UserIdentity       string `gorm:"type:varchar(36);"`
	ParentID           int    `gorm:"type:int"`
	RepositoryIdentity string `gorm:"type:varchar(36);not null"`
	Ext                string `gorm:"type:varchar(255);not null;comment:'文件或文件夹类型'"`
	Name               string `gorm:"type:varchar(255);"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
