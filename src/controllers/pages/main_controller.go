package pagecontrollers

import "github.com/gofiber/fiber/v2"

func InitMainController() *fiber.App {
	app := fiber.New()

	app.Get("/home-page", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("main/home-page", fiber.Map{
			"Title": "Hello, World!",
		}, "layout/main-layout")
	})

	app.Get("/login-page", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("main/login-page", fiber.Map{
			"Title": "Hello, World!",
		}, "layout/main-layout")
	})

	return app
}
