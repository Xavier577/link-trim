package database

import (
	"example/trim-server/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  global.Env("DB_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
