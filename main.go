package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server Starting")
	})

	app.Get("/tegar", func(c *fiber.Ctx) error {
		return c.SendString("halo saya tegar")
	})

	app.Listen(":3000")
}
