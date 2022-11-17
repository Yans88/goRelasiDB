package migrations

import (
	"fmt"
	"github.com/yansen/goRelasiDB/database"
	"github.com/yansen/goRelasiDB/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println("can't running migration")
	}
}
