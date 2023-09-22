package database

import (
	"fmt"
	"go_mini-project/config"
	"go_mini-project/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//db using gorm

var DB *gorm.DB

var (
	DB_USERNAME string = config.GetConfig("DB_USERNAME")
	DB_PASSWORD string = config.GetConfig("DB_PASSWORD")
	DB_NAME     string = config.GetConfig("DB_NAME")
	DB_HOST     string = config.GetConfig("DB_HOST")
	DB_PORT     string = config.GetConfig("DB_PORT")
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("ERROR SAAT MENGHUBUNGKAN KE DATABASE: %s", err)
	}

	log.Println("Terhubung ke MySQL")

	DB.AutoMigrate(&model.Tiket{}, &model.Kereta{}, &model.Pemesan{})

}
