package test

import (
	"cloud-disk/core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := "root:565656@tcp(localhost:3306)/cloud-disk?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatal(err)
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = tx.AutoMigrate(&models.UserBasic{})
	if err != nil {
		t.Fatal(err)
	}
	tx.Commit()
}
