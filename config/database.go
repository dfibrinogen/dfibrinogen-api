package config

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func InitDB(databaseUrl string) *gorm.DB {

	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// Enable Log Mode if you want to enable Log on Database Query ( Gorm )
	// And disable Log Mode if you want to deploy to Production
	if IsProduction {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	db.AutoMigrate(
		&model.DUser{},
		&model.DProfile{},
		&model.DCategory{},
		&model.DPost{},
		&model.DPostComment{},
		&model.DPostLike{},
	)

	return db
}
