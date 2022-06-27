package configs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() (result *gorm.DB) {
	configs := GetConfig()
	var err error
	if db == nil {
		db, err = gorm.Open(postgres.Open(configs.PgUri), &gorm.Config{})
		if err != nil {
			log.Fatal("Can't connect to db")
			os.Exit(1)
		}
	}
	result = db
	return db
}
