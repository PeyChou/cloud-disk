package models

import "gorm.io/gorm"

type ShareBasic struct {
	gorm.Model
	Identity               string `gorm:"size:36"`
	UserIdentity           string `gorm:"size:36"`
	RepositoryIdentity     string `gorm:"size:36;not null"`
	UserRepositoryIdentity string `gorm:"size:36;not null"`
	ExpireTime             int    `gorm:"default:0;comment:失效时间(s)，0表示永不失效"`
	ClickNum               int    `gorm:"default:0;comment:点击次数"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
