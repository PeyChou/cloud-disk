package models

import "gorm.io/gorm"

type RepositoryPool struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36);"`
	Hash     string `gorm:"type:varchar(32);comment:'文件唯一标识'"`
	Name     string `gorm:"type:varchar(255);"`
	Ext      string `gorm:"type:varchar(30);comment:'文件拓展名'"`
	Size     int64  `gorm:"type:bigint;comment:'文件大小'"`
	Path     string `gorm:"type:varchar(255);comment:'文件路径'"`
}

// TableName 设置表名为 "repository_pool"
func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
