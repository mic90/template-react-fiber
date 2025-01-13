package main

import "github.com/gofiber/fiber/v2"

func configureRouter(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
