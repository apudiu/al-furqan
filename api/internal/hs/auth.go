package hs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetToken(c *fiber.Ctx) (token *jwt.Token, ok bool) {
	userInCtx := c.Locals("user")
	if userInCtx == nil {
		return
	}

	ok = true
	token = userInCtx.(*jwt.Token)
	return
}

type TokenPayload struct {
	// Token subject like, user id
	Sub   uint    `json:"sub"`
	Email string  `json:"email"`
	Exp   float64 `json:"exp"`
}

func (tp *TokenPayload) ParseFromCtx(ctx *fiber.Ctx) bool {
	token, ok := GetToken(ctx)
	if !ok {
		return ok
	}

	ok = tp.ParseFromJwt(token)
	return ok
}

func (tp *TokenPayload) ParseFromJwt(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)

	ok := tp.ParseFromClaims(claims)
	return ok
}

func (tp *TokenPayload) ParseFromClaims(claims jwt.MapClaims) bool {
	// extract sub
	sub, subFound := claims["sub"]
	if !subFound {
		return false
	}

	subInt, ok := sub.(float64) // all int values are of type float64 in claims
	if !ok {
		return false
	}
	tp.Sub = uint(subInt)

	// extract email
	email, emailFound := claims["email"]
	if !emailFound {
		return false
	}
	tp.Email = email.(string)

	// extract ext
	exp, expFound := claims["exp"]
	if !expFound {
		return false
	}
	tp.Exp = exp.(float64)

	return true
}

func (tp *TokenPayload) ToJwtMapClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"sub":   tp.Sub,
		"email": tp.Email,
		"exp":   tp.Exp,
	}
}

func NewTokenPayload(sub uint, email string, exp float64) TokenPayload {
	tp := TokenPayload{
		Sub:   sub,
		Email: email,
		Exp:   exp,
	}
	return tp
}

func NewEmptyTokenPayload() TokenPayload {
	tp := TokenPayload{}
	return tp
}
