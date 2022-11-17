package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/database/migrations"
	"github.com/yansen/goRelasiDB/routes"
)

func main() {
	database.DBInit()
	migrations.Migration()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hallo"})
	})

	routes.RouteInit(app)
	app.Listen(":8000")
}
