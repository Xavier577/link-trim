package middlewares

import (
	"example/trim-server/global"
	stringParser "example/trim-server/shared/string-parser"
	"example/trim-server/shared/tokenizer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		access_token := tokenizer.ExtractTokenFromHeader(context)

		if stringParser.IsStringNull(access_token) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UNAUTHORIZED"})
			return
		}

		isValid, claim, err := tokenizer.VerifyToken(tokenizer.TokenVerifyOptions{
			Secret:      global.Env("JWT_SECRET"),
			SignedToken: access_token,
		})

		if err != nil && err.Error() == "TOKEN_EXPIRED" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "TOKEN_EXPIRED"})
			return
		}

		if err != nil && err.Error() == "TOKEN_PARSE_ERR" {
			log.Fatal(err.Error())
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR"})
			return
		}

		if !isValid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UNAUTHORIZED"})
			return
		}

		context.Set("user_id", claim.Payload["user_id"])
		context.Next()

	}
}
