package authhandler

import (
	"errors"
	"github.com/apudiu/alfurqan/internal/hs"
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
		return c.Status(fiber.StatusBadRequest).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	if !isEmail(input.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   errors.New("invalid email address"),
		}))
	}

	user, token, err := authservice.SignIn(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   errors.New("invalid email address"),
		}))
	}

	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		D: fiber.Map{
			"user":  user,
			"token": token,
		},
	}))
}

func SignUp(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	token, err := authservice.SignUp(user)
	if err != nil {
		return c.Status(500).JSON(hs.Res(hs.ResData{
			Msg: "User creation failed",
			D:   err,
		}))
	}

	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Successfully registered",
		D: fiber.Map{
			"user":  user,
			"token": token,
		},
	}))
}

func AuthUser(c *fiber.Ctx) error {
	token, ok := hs.GetToken(c)
	if !ok {
		return errors.New("failed to parse token")
	}

	user, err := authservice.GetAuthUser(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(hs.Res(hs.ResData{
			D: err,
		}))
	}

	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		D:      user,
	}))
}

func SignOut(c *fiber.Ctx) error {
	token, ok := hs.GetToken(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
		}))
	}

	err := authservice.SignOut(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(hs.Res(hs.ResData{
			Msg: "Error signing out",
			D:   err,
		}))
	}

	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Signed Out",
	}))
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
