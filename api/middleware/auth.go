package middleware

import (
	"github.com/apudiu/alfurqan/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.Config("KEY")),
		},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"msg":    "Missing or malformed token",
			"d":      nil,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "error",
		"msg":    "Invalid or expired token",
		"d":      nil,
	})
}
