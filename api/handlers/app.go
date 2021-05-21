package handlers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize() {
	dsn := "root:LinuxKU13!@tcp(docker.for.mac.localhost:3306)/intranet_v2?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = db
}
