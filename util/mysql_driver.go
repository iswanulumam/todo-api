package util

import (
	"fmt"
	"todo-layered/config"
	"todo-layered/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDriver(config *config.AppConfig) *gorm.DB {
	var uri string

	uri = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	DatabaseMigration(db)

	return db
}

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(entities.Todo{})
}
