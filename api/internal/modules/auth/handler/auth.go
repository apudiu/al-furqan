package authhandler

import (
	"github.com/apudiu/alfurqan/internal/model"
	authservice "github.com/apudiu/alfurqan/internal/modules/auth/service"
	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	return nil
}

func SignUp(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": false,
			"msg":    "Review your input",
			"d":      err,
		})
	}

	err := authservice.SignUp(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": false,
			"msg":    "User creation failed",
			"d":      err,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "Successfully registered.",
		"d":      user,
	})
}

func SignOut(c *fiber.Ctx) error {
	return nil
}

func RequestResetPassword(c *fiber.Ctx) error {
	return nil
}

func ResetPassword(c *fiber.Ctx) error {
	return nil
}
