package routes

import (
	helper "kotuSozApi/helper"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/check", helper.KotuSozCheck)
	app.Post("/api/filter", helper.KotuSozFilter)
}
