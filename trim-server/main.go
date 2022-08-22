package main

import (
	"example/trim-server/config"
	"example/trim-server/database"
	"example/trim-server/global"
	"example/trim-server/shared/orm"

	"github.com/gin-gonic/gin"

	_ "example/trim-server/docs"
)

func init() {
	orm.LoadDriver(database.Connect())

}

// @title Trim-Sever
// @version 1.0
// @description This is a simple link shortner api. You can visit the GitHub repository at https://github.com/Xavier577/link-trim

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
// @query.collection.format multi

//@securityDefinitions.apikey Authorization
//@in header
//@name Authorization

func main() {
	router := gin.Default()
	config.RoutesConfig(router)

	router.Run(global.Env("ADDRESS"))
}
