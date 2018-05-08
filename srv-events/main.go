package main

import (
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/db"
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/handler"
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/repository"
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

	handler.NewEventHandler(v1, repository.InitEventRepo(database))

	e.Logger.Fatal(e.Start(":5003"))
}
