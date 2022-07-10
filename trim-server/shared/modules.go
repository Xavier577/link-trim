package shared

import (
	"example/trim-server/auth"
	trimmedlinks "example/trim-server/trimmed-links"
	"example/trim-server/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoutesModule(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		auth.Routes(v1)
		users.Routes(v1)
		trimmedlinks.Routes(v1)
	}
}

func RepositoryModule(db_driver *gorm.DB) {
	users.InitializeUserRepository(db_driver)
	trimmedlinks.InitializeLinkRepository(db_driver)
}
