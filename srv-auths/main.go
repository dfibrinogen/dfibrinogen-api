package main

import (
	"log"
	"runtime"
	"github.com/labstack/echo"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/db"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/handler"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/repository"
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

	handler.NewAuthHandler(e, repository.InitAuthRepo(database))

	e.Logger.Fatal(e.Start(":5001"))
}
