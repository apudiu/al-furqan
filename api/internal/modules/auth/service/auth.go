package authservice

import (
	"github.com/apudiu/alfurqan/internal/model"
	"github.com/gofiber/fiber/v2"
)

func SignIn(username, password string) (model.User, error) {

}

func SignUp(c *fiber.Ctx) error {
	//
}

func SignOut(c *fiber.Ctx) error {
	//
}

func RequestResetPassword(c *fiber.Ctx) error {
	//
}

func ResetPassword(c *fiber.Ctx) error {
	//
}
