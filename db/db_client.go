package db

import (
	"github.com/dfibrinogen/dfibrinogen-api/module/auths"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func InitDatabase(databaseURL string) *gorm.DB {

	dbGorm, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	//defer dbGorm.Close()

	dbGorm.LogMode(true)

	dbGorm.AutoMigrate(
		&auths.Auth{},
	)

	return dbGorm
}
