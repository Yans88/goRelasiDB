package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yansen/goRelasiDB/controllers"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)

	app.Get("/locker", controllers.LockerGetAll)
	app.Post("/locker", controllers.CreateLocker)

	app.Get("/posts", controllers.PostGetAll)
	app.Post("/posts", controllers.CreatePost)

	app.Get("/tags", controllers.TagGetAll)
	app.Post("/tags", controllers.CreateTag)
}
