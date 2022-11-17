package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/models"
)

func PostGetAll(c *fiber.Ctx) error {
	var posts []models.Post
	database.DB.Preload("User").Find(&posts)
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"message": "can't handle request",
			"data":    err,
		})
	}
	if post.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, title is required",
			"data":    "",
		})
	}
	if post.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, user_id is required",
			"data":    "",
		})
	}

	if len(post.TagsID) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request, tags_id is required",
			"data":    "",
		})
	}
	database.DB.Create(&post)
	for _, tagID := range post.TagsID {
		postTag := new(models.PostTag)
		postTag.PostID = post.ID
		postTag.TagID = tagID
		database.DB.Create(&postTag)
	}
	return c.JSON(fiber.Map{
		"message": "ok",
		"data":    post,
	})
}
