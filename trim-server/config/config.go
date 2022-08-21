package config

import (
	"example/trim-server/auth"
	trimmedlinks "example/trim-server/trimmed-links"
	"example/trim-server/users"
	"net/http"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesConfig(r *gin.Engine) {

	// use ginSwagger middleware to serve the API docs
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "docs/swagger/index.html")
	})

	{
		r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth.Routes(r.Group("/auth"))
		users.Routes(r.Group("/user"))
		trimmedlinks.Routes(r.Group("/trimmed-link"))
	}
}
