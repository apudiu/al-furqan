package authhandler

import (
	"errors"
	"github.com/apudiu/alfurqan/internal/hs"
	"github.com/apudiu/alfurqan/internal/model"
	authservice "github.com/apudiu/alfurqan/internal/modules/auth/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
)

func SignIn(c echo.Context) error {
	type signInInput struct {
		Email, Password string
	}

	input := new(signInInput)
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	if !isEmail(input.Email) {
		return c.JSON(http.StatusBadRequest, hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   errors.New("invalid email address"),
		}))
	}

	user, token, err := authservice.SignIn(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   errors.New("invalid email address"),
		}))
	}

	return c.JSON(http.StatusOK, hs.Res(hs.ResData{
		Status: true,
		D: hs.AnyMap{
			"user":  user,
			"token": token,
		},
	}))
}

func SignUp(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	token, err := authservice.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, hs.Res(hs.ResData{
			Msg: "User creation failed",
			D:   err,
		}))
	}

	return c.JSON(http.StatusOK, hs.Res(hs.ResData{
		Status: true,
		Msg:    "Successfully registered",
		D: hs.AnyMap{
			"user":  user,
			"token": token,
		},
	}))
}

func AuthUser(c echo.Context) error {
	token, ok := hs.GetToken(c)
	if !ok {
		return errors.New("failed to parse token")
	}

	user, err := authservice.GetAuthUser(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, hs.Res(hs.ResData{
			D: err,
		}))
	}

	return c.JSON(http.StatusOK, hs.Res(hs.ResData{
		Status: true,
		D:      user,
	}))
}

func SignOut(c echo.Context) error {
	token, ok := hs.GetToken(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, hs.Res(hs.ResData{
			Msg: "Review your input",
		}))
	}

	err := authservice.SignOut(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, hs.Res(hs.ResData{
			Msg: "Error signing out",
			D:   err,
		}))
	}

	return c.JSON(http.StatusOK, hs.Res(hs.ResData{
		Status: true,
		Msg:    "Signed Out",
	}))
}

func RequestResetPassword(c echo.Context) error {
	return nil
}

func ResetPassword(c echo.Context) error {
	return nil
}

// helpers

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
