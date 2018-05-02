package main

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/db"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/service"
	"github.com/labstack/echo"
	"log"
	"runtime"
)

func init() {

	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	databaseURL := "host=localhost user=davian dbname=db_dfibrinogen_api sslmode=disable password="

	database := db.InitDatabase(databaseURL)

	e := echo.New()

	v1 := e.Group("/api/v1")

	service.NewEventService(v1, repository.InitEventRepo(database))

	e.Logger.Fatal(e.Start(":5004"))
}
