package middlewares

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtProtect() fiber.Handler {
	return jwtware.New(jwtware.Config{
		TokenLookup: "cookie:jwtToken", //we can set this by cookie or authorization header
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})
}
