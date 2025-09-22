package migration

import (
	"log"
	"ms-golang-echo/config"
	"ms-golang-echo/internal/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		panic("[SYSTEM ERROR] Migration failed: " + err.Error())
	}

	log.Println("[SYSTEM] Migration completed")
}
