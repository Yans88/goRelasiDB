package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/models"
)

func LockerGetAll(c *fiber.Ctx) error {
	var lockers []models.Locker
	database.DB.Preload("User").Find(&lockers)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    lockers,
	})
}

func CreateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)
	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"message": "can't handle request",
			"data":    "",
		})
	}
	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, code is required",
			"data":    "",
		})
	}
	if locker.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, user_id is required",
			"data":    "",
		})
	}
	database.DB.Create(&locker)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    locker,
	})
}
