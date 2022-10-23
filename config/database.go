package config

import (
	"gobook/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/gobook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Book{}, &models.User{})

	return db
}
