package main

import (
	"fmt"
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/router"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	//todo: parse .env file
	isLocal := os.Getenv("DEBUG") == "1"
	e.Debug = isLocal

	// connect to DB
	database.ConnectDB()

	// home route (public)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Alhum-du-lillah")
	})

	// setup router
	router.SetupRoutes(e)

	// print routes
	if isLocal {
		printRoutes(e.Routes())
	}

	log.Fatalln(
		e.Start(":3001"),
	)
}

func printRoutes(rs []*echo.Route) {
	for _, r := range rs {
		fmt.Printf("%+6s %s\n", r.Method, r.Path)
	}
}
