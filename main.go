package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Champion struct {
	Name     string `json:"name"`
	Weapon   string `json:"weapon"`
	Hometown string `json:"hometown"`
}

var champion Champion

func createChampion(ctx *fiber.Ctx) error {

	body := new(Champion)

	err := ctx.BodyParser(body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	champion = Champion{
		Name:     body.Name,
		Weapon:   body.Weapon,
		Hometown: body.Hometown,
	}

	return ctx.Status(fiber.StatusOK).JSON(champion)
}

func getChampion(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(champion)
}

func main() {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/champion", getChampion)
	app.Post("/champion", createChampion)

	app.Listen(":8080")
}
