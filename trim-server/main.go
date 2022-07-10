package main

import (
	"example/trim-server/config"
	"example/trim-server/database"
	"example/trim-server/shared"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	shared.RoutesModule(router)

	db := database.Connect()

	shared.RepositoryModule(db)

	router.Run(config.Env("ADDRESS"))
}
