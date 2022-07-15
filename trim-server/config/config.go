package config

import (
	"example/trim-server/auth"
	trimmedlinks "example/trim-server/trimmed-links"
	"example/trim-server/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesConfig(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		auth.Routes(v1.Group("/auth"))
		users.Routes(v1.Group("/user"))
		trimmedlinks.Routes(v1.Group("/trimmed-link"))
	}
}

func RepositoryInitializer(db_driver *gorm.DB) {
	users.InitializeUserRepository(db_driver)
	trimmedlinks.InitializeLinkRepository(db_driver)
}
