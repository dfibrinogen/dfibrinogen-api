package main

import (
	"runtime"
	"log"
	"github.com/labstack/echo"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/service"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
)

func init() {

	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {


	e := echo.New()

	v1 := e.Group("/api/v1")

	service.NewEventService(v1, repository.InitEventRepo())

	e.Logger.Fatal(e.Start(":5004"))
}
