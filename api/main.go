// main.go
package main

import (
	"auth-api/database"
	"auth-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// GORMセット
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":80")
}
