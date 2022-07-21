package config

import (
	"example/trim-server/auth"
	trimmedlinks "example/trim-server/trimmed-links"
	"example/trim-server/users"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesConfig(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth.Routes(v1.Group("/auth"))
		users.Routes(v1.Group("/user"))
		trimmedlinks.Routes(v1.Group("/trimmed-link"), router)
	}
}
