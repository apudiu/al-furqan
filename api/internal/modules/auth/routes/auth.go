package authroutes

import (
	authhandler "github.com/apudiu/alfurqan/internal/modules/auth/handler"
	"github.com/apudiu/alfurqan/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	auth := router.Group("auth")

	auth.Get("/me", middleware.Protected(), authhandler.AuthUser)
	auth.Post("/sign-in", authhandler.SignIn)
	auth.Post("/sign-up", authhandler.SignUp)
	auth.Post("/sign-out", middleware.Protected(), authhandler.SignOut)
	auth.Post("/reset-request", authhandler.RequestResetPassword)
	auth.Patch("/reset-password", authhandler.ResetPassword)
}
