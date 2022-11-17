package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/models"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Preload("Locker").Preload("Posts").Find(&users)
	return c.JSON(fiber.Map{
		"data":    users,
		"message": "ok",
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"message": "can't handle request",
			"data":    "",
		})
	}
	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, name is required",
			"data":    "",
		})
	}
	database.DB.Create(&user)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    user,
	})
}
