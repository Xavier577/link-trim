package database

import (
	"example/trim-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Env("DB_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
