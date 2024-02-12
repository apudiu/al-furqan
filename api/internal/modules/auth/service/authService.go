package authservice

import (
	"errors"
	"github.com/apudiu/alfurqan/config"
	"github.com/apudiu/alfurqan/internal/helpers"
	"github.com/apudiu/alfurqan/internal/model"
	userservice "github.com/apudiu/alfurqan/internal/modules/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func SignIn(email, password string) (user model.User, signedToken string, err error) {
	// check if user exist with email
	user, err = userservice.GetByEmail(email)
	if err != nil {
		err = errors.New("invalid credential")
		return
	}

	// verify password
	if !checkPasswordHash(password, user.Password) {
		err = errors.New("invalid credential")
		return
	}

	// make token
	signedToken, err = makeUserToken(&user)
	if err != nil {
		err = errors.New("token creation err")
	}

	// save token hash to db
	ok := userservice.UpdateTokenHash(&user, signedToken)
	if !ok {
		err = errors.New("token update err")
	}

	return
}

func SignUp(u *model.User) (signedToken string, err error) {
	// check if user doesn't exist with email
	_, err = userservice.GetByEmail(u.Email)
	if err == nil {
		err = errors.New(u.Email + " already taken")
		return
	}

	// create user

	pwHash, err := helpers.HashStr(u.Password, 14)
	if err != nil {
		return
	}
	u.Password = pwHash

	err = userservice.Create(u)
	if err != nil {
		return
	}

	// make token
	signedToken, err = makeUserToken(u)
	if err != nil {
		return
	}

	// save token hash to db
	ok := userservice.UpdateTokenHash(u, signedToken)
	if !ok {
		err = errors.New("token update err")
	}
	return
}

func SignOut(t *jwt.Token) error {
	tp := helpers.NewEmptyTokenPayload()
	ok := tp.ParseFromJwt(t)
	if !ok {
		return errors.New("failed to parse token")
	}

	// delete saved token
	user, err := userservice.GetById(tp.Sub)
	if err != nil {
		return err
	}

	ok = userservice.UpdateTokenHash(&user, "")
	if !ok {
		return errors.New("failed to reset token")
	}

	return nil
}

func GetAuthUser(t *jwt.Token) (user model.User, err error) {
	tp := helpers.NewEmptyTokenPayload()
	ok := tp.ParseFromJwt(t)

	if !ok {
		err = errors.New("invalid token")
	}

	// retrieve the user
	user, err = userservice.GetById(tp.Sub)
	return
}

func RequestResetPassword(c *fiber.Ctx) error {
	return nil
}

func ResetPassword(c *fiber.Ctx) error {
	return nil
}

// helpers

// checkPasswordHash compare password with hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func makeUserToken(u *model.User) (token string, err error) {
	expiry := time.Now().Add(time.Hour * 72).Unix()

	tokenPayload := helpers.NewTokenPayload(u.ID, u.Email, float64(expiry))
	claims := tokenPayload.ToJwtMapClaims()

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = rawToken.SignedString([]byte(config.Config("KEY")))
	return
}
