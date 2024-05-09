package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
	"github.com/lpernett/godotenv"
	apicontrollers "github.com/rivandra0/fiber-mvc-template/src/controllers/apis"
	pagecontrollers "github.com/rivandra0/fiber-mvc-template/src/controllers/pages"

	"github.com/rivandra0/fiber-mvc-template/src/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./src/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.InitDatabase()

	app.Use(helmet.New())

	app.Static("/static", "./public")

	app.Mount("/api", apicontrollers.InitApiRouter())
	app.Mount("/", pagecontrollers.InitMainController())
	app.Mount("/dashboard", pagecontrollers.InitDashboardController())

	if err := app.Listen(":3000"); err != nil {
		fmt.Print("error occur:", err)
		return
	}
}
