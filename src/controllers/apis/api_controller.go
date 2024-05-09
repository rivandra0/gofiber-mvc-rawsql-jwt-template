package apicontrollers

import (
	"github.com/gofiber/fiber/v2"
)

func InitApiRouter() *fiber.App {

	app := fiber.New()
	app.Mount("/auth", InitAuthController())
	app.Mount("/users", InitUserController())
	return app
}
