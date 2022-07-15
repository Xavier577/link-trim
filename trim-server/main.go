package main

import (
	"example/trim-server/config"
	"example/trim-server/database"
	"example/trim-server/global"

	"github.com/gin-gonic/gin"
)

func init() {
	db := database.Connect()
	config.RepositoryInitializer(db)
}

func main() {
	router := gin.Default()
	config.RoutesConfig(router)
	router.Run(global.Env("ADDRESS"))
}
