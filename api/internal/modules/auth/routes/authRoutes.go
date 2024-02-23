package authroutes

import (
	authhandler "github.com/apudiu/alfurqan/internal/modules/auth/handler"
	"github.com/labstack/echo/v4"
)

func SetupPublic(router *echo.Group) {
	auth := router.Group("/auth")

	auth.POST("/sign-in", authhandler.SignIn)
	auth.POST("/sign-up", authhandler.SignUp)
	auth.POST("/reset-request", authhandler.RequestResetPassword)
	auth.PATCH("/reset-password", authhandler.ResetPassword)
}

func SetupPrivate(router *echo.Group) {
	auth := router.Group("/auth")

	auth.GET("/me", authhandler.AuthUser)
	auth.POST("/sign-out", authhandler.SignOut)
}
