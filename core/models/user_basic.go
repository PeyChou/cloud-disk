package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36);not null;"`
	Name     string `gorm:"type:varchar(16);not null;"`
	Password string `gorm:"type:varchar(32);not null;"`
	Email    string `gorm:"type:varchar(32);not null;"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
