package pagecontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rivandra0/fiber-mvc-template/src/middlewares"
)

func InitDashboardController() *fiber.App {
	app := fiber.New()
	middlewares.JwtProtect()

	JwtProtect := middlewares.JwtProtect()

	app.Get("/", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("dashboard/main-dashboard-page", fiber.Map{
			"Title": "Hello, World!",
		}, "layout/dashboard-layout")
	})

	app.Get("/main-dashboard-page", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("dashboard/main-dashboard-page", fiber.Map{
			"Title": "Hello, World!",
		}, "layout/dashboard-layout")
	})

	app.Get("/user-dashboard-page", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("dashboard/user-dashboard-page", fiber.Map{
			"Title": "Hello, World!",
		}, "layout/dashboard-layout")
	})

	return app
}
