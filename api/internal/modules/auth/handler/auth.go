package authhandler

import (
	"errors"
	"fmt"
	"github.com/apudiu/alfurqan/internal/helpers"
	"github.com/apudiu/alfurqan/internal/model"
	authservice "github.com/apudiu/alfurqan/internal/modules/auth/service"
	"github.com/gofiber/fiber/v2"
	"net/mail"
)

func SignIn(c *fiber.Ctx) error {
	type signInInput struct {
		Email, Password string
	}

	input := new(signInInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    "Review your input",
			"d":      err,
		})
	}

	if !isEmail(input.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    "Review your input",
			"d":      errors.New("invalid email address"),
		})
	}

	user, token, err := authservice.SignIn(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    "Review your input",
			"d":      errors.New("invalid email address"),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "",
		"d": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
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

	token, err := authservice.SignUp(user)
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
		"d": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}

func AuthUser(c *fiber.Ctx) error {
	token, ok := helpers.GetToken(c)
	if !ok {
		return errors.New("failed to parse token")
	}

	user, err := authservice.GetAuthUser(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"msg":    "Review your input",
			"d":      err,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "",
		"d":      user,
	})
}

func SignOut(c *fiber.Ctx) error {
	fmt.Println("IMPLEMENT SIGN OUT")
	return nil
}

func RequestResetPassword(c *fiber.Ctx) error {
	return nil
}

func ResetPassword(c *fiber.Ctx) error {
	return nil
}

// helpers

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
