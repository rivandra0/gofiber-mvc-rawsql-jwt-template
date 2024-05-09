package apicontrollers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rivandra0/fiber-mvc-template/src/models"
	"github.com/rivandra0/fiber-mvc-template/src/queries"
)

var userQuery queries.UserQuery = queries.UserQuery{}

func InitAuthController() *fiber.App {
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {

		bodyUser := new(models.User)
		if err := c.BodyParser(&bodyUser); err != nil {
			fmt.Println("error at parsing")
			return err
		}

		userFromDb, err := userQuery.GetOne(bodyUser.Email, bodyUser.Password)
		if err != nil {
			c.JSON(fiber.Map{"Message": err.Error()})
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		fmt.Print(userFromDb.Email)
		fmt.Print("name is", userFromDb.UserName, "and Password is", userFromDb.Password)

		// Create the Claims
		claims := jwt.MapClaims{
			"Id":               userFromDb.Id,
			"Email":            userFromDb.Email,
			"userFromDbName":   userFromDb.UserName,
			"GroupId":          userFromDb.GroupId,
			"Tier":             userFromDb.Tier,
			"IsEmailConfirmed": userFromDb.IsEmailConfirmed,
			"IsActive":         userFromDb.IsActive,
			"exp":              time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			fmt.Print(err.Error())
			c.JSON(fiber.Map{"Message": err.Error()})
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// fmt.Print("token", t)
		cookie := new(fiber.Cookie)
		cookie.Name = "jwtToken"
		cookie.Value = t
		fmt.Println(cookie.Name)
		c.Cookie(cookie)
		c.SendStatus(fiber.StatusOK)
		// return c.Redirect("dashboard/main-dashboard-page")
		return c.JSON(fiber.Map{"message": "Successfully logged in"})
	})

	return app
}
