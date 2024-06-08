package internal

import (
	"log"

	"github.com/akposiyefa/go-gin-migration/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Envs.DB_URL), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Sorry unable to connect to database", err)
		return
	}
}
