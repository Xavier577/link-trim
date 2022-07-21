package database

import (
	"example/trim-server/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  global.Env("DB_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}

	LoadMigrations(db)

	return db
}
