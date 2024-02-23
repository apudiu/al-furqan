package router

import (
	"github.com/apudiu/alfurqan/config"
	authRoutes "github.com/apudiu/alfurqan/internal/modules/auth/routes"
	userRoutes "github.com/apudiu/alfurqan/internal/modules/user/routes"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// versioning
	apiVer := e.Group("/v1")

	// Register public routes
	authRoutes.SetupPublic(apiVer)
	userRoutes.SetupPublic(apiVer)

	// Register private routes

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.Config("KEY")),
	}))

	authRoutes.SetupPrivate(apiVer)
	userRoutes.SetupPrivate(apiVer)
}
