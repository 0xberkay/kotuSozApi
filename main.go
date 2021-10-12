package main

import (
	"kotuSozApi/database"
	"kotuSozApi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.Setup(app)

	database.Connect()

	app.Listen(":3000")
}
