package apicontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rivandra0/fiber-mvc-template/src/middlewares"
)

func InitUserController() *fiber.App {

	app := fiber.New()

	JwtProtect := middlewares.JwtProtect()
	//getOne
	app.Get("/:id", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.SendString("GetOne User: " + c.Params("id"))
	})

	//getMany
	app.Get("/", JwtProtect, func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		_ = name
		// Render index within layouts/main
		return c.SendString("GetMany User")
	})

	//insertOne
	app.Post("/", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.SendString("PostOne User")
	})

	//updateOne
	app.Put("/", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.SendString("PutOne User")
	})

	//deleteOne
	app.Delete("/", JwtProtect, func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.SendString("DeleteOne User")
	})

	return app
}
