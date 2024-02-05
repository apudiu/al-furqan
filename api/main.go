package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Alhum-du-lillah")
	})

	log.Fatalln(
		app.Listen(":3001"),
	)
}
