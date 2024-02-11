package authservice

import (
	"errors"
	"github.com/apudiu/alfurqan/config"
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/internal/model"
	userservice "github.com/apudiu/alfurqan/internal/modules/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	signedToken, err = token.SignedString([]byte(config.Config("KEY")))
	if err != nil {
		err = errors.New("token creation err")
	}

	return
}

func SignUp(u *model.User) error {
	// check if user doesn't exist with email
	_, err := userservice.GetByEmail(u.Email)
	if err == nil {
		return errors.New(u.Email + " already taken")
	}

	// create user

	pwHash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwHash

	err = userservice.Create(u)
	return err
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

// helpers

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Email == "" {
		return false
	}
	if !checkPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// checkPasswordHash compare password with hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
