package authservice

import (
	"errors"
	"github.com/apudiu/alfurqan/config"
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

	return
}

func SignUp(u *model.User) (token string, err error) {
	// check if user doesn't exist with email
	_, err = userservice.GetByEmail(u.Email)
	if err == nil {
		err = errors.New(u.Email + " already taken")
		return
	}

	// create user

	pwHash, err := hashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = pwHash

	err = userservice.Create(u)
	if err != nil {
		return
	}

	// make token
	token, err = makeUserToken(u)
	return
}

func SignOut(c *fiber.Ctx) error {
	return nil
}

func GetAuthUser(t *jwt.Token) (user model.User, err error) {
	claims := t.Claims.(jwt.MapClaims)

	id, found := claims["id"]
	if !found {
		err = errors.New("invalid user")
		return
	}

	userId := id.(float64)

	// retrieve the user
	user, err = userservice.GetById(uint(userId))

	return
}

func RequestResetPassword(c *fiber.Ctx) error {
	return nil
}

func ResetPassword(c *fiber.Ctx) error {
	return nil
}

// helpers

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compare password with hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func makeUserToken(u *model.User) (token string, err error) {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"ext":   time.Now().Add(time.Hour * 72).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = rawToken.SignedString([]byte(config.Config("KEY")))
	return
}
