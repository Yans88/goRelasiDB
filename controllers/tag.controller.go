package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/models"
)

func TagGetAll(c *fiber.Ctx) error {
	var tags []models.TagResponseWithPost
	database.DB.Preload("Posts").Find(&tags)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    tags,
	})
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)
	if err := c.BodyParser(tag); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"message": "can't handle request",
			"data":    err,
		})
	}
	if tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, name is required",
			"data":    "",
		})
	}

	database.DB.Create(&tag)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    tag,
	})
}
