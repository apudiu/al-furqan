package main

import (
	"encoding/json"
	"fmt"
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/router"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	// connect to DB
	database.ConnectDB()

	// home route (public)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Alhum-du-lillah")
	})

	// setup router
	router.SetupRoutes(e)

	// print routes
	rs := e.Routes()
	rj, _ := json.MarshalIndent(rs, "", "  ")
	fmt.Println(string(rj))

	log.Fatalln(
		e.Start(":3001"),
	)
}
