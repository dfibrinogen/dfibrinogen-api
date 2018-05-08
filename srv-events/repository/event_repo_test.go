package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
	"time"
)

func TestEventRepository_FetchEventAll(t *testing.T) {

	// Initialize Mock Database
	db, m := mockUpDatabase(t)

	// Prepare Mock Rows
	fields := []string{"id", "name", "location", "created_at", "updated_at", "deleted_at"}
	rows := sqlmock.NewRows(fields)
	rows = rows.AddRow("01", "Name 1", "Location 1", time.Now(), time.Now(), nil)
	rows = rows.AddRow("02", "Name 2", "Location 2", time.Now(), time.Now(), nil)

	//m.ExpectQuery("^SELECT (.+) FROM \"events\" WHERE \"events\".\"deleted_at\" IS NULL$").
	m.ExpectQuery(fixedFullRe("SELECT * FROM \"events\" WHERE \"events\".\"deleted_at\" IS NULL")).
		WillReturnRows(rows)

	repo := InitEventRepo(db)

	list, err := repo.FetchEventAll()

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestEventRepository_FetchEventByID(t *testing.T) {

	//// Initialize Mock Database
	//db, m := mockUpDatabase(t)
	//
	//// Prepare Mock Rows
	//fields := []string{"id", "name", "location", "created_at", "updated_at", "deleted_at"}
	//rows := sqlmock.NewRows(fields)
	//rows = rows.AddRow("01", "Name 1", "Location 1", time.Now(), time.Now(), nil)
	//
	//m.ExpectQuery(fixedFullRe("SELECT * FROM \"events\" WHERE \"events\".\"deleted_at\" IS NULL AND ((\"events\".\"id\" = $1)) ORDER BY \"events\".\"id\" ASC LIMIT 1")).
	//	WillReturnRows(rows)
	//
	//repo := InitEventRepo(db)
	//
	//data, err := repo.FetchEventByID("01")
	//
	//assert.NoError(t, err)
	//assert.Len(t, data, 1)
}

func fixedFullRe(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

func mockUpDatabase(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {

	db, m, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a SQL database connection", err)
	}

	dbGorm, errGorm := gorm.Open("postgres", db)
	if errGorm != nil {
		t.Fatalf("an error '%s' was not expected when opening a GORM database connection", err)
	}

	return dbGorm, m
}
