package helper

import (
	"kotuSozApi/database"
	"kotuSozApi/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var kotusoz models.Tbl

// Check Helper
func KotuSozCheck(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	str := data["str"]
	split := strings.Split(str, " ")
	loopStreak := len(split)

	cout := 0
	for i := 0; i < loopStreak; i++ {
		if split[i] != " " && split[i] != "" && split[i] != "  " {
			database.DB.Table("tlbs").Find(&kotusoz, "kotusoz = ?", split[i])
			if kotusoz.Kotusoz != "" {
				cout++
			} else {
				if i+1 < loopStreak {
					multiString := split[i] + " " + split[i+1]

					database.DB.Find(&kotusoz, "kotusoz = ?", multiString)
				}
				if kotusoz.Kotusoz != "" {
					cout++
				}
			}
		}

		kotusoz = models.Tbl{}
	}
	// Return
	if cout < 0 {
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kotu soz bulunamadı", "count": cout})
	} else {
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kotu soz bulundu", "count": cout})
	}

}

// Filter Helper
func KotuSozFilter(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	str := data["str"]
	filterValue := data["filter"]

	split := strings.Split(str, " ")
	loopStreak := len(split)

	for i := 0; i < loopStreak; i++ {
		if split[i] != " " && split[i] != "" && split[i] != "  " {
			database.DB.Find(&kotusoz, "kotusoz = ?", split[i])
			if kotusoz.Kotusoz != "" {
				split[i] = filterValue
			} else {
				if i+1 < loopStreak {
					multiString := split[i] + " " + split[i+1]

					database.DB.Find(&kotusoz, "kotusoz = ?", multiString)
				}
				if kotusoz.Kotusoz != "" {
					split[i] = filterValue
					split[i+1] = filterValue
				}
			}
		}

		kotusoz = models.Tbl{}
	}

	connectedString := strings.Join(split, " ")

	// Return
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": connectedString})

}
