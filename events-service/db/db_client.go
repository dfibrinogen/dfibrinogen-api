package db

import (
	"database/sql"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func InitDatabase(databaseURL string) *sql.DB {

	dbGorm, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer dbGorm.Close()

	dbGorm.LogMode(true)

	dbGorm.AutoMigrate(
		&model.Event{},
	)

	dbNative, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer dbNative.Close()

	return dbNative
}
