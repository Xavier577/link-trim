package tokenizer

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractTokenFromHeader(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")

	if bearerTokenChars := strings.Split(bearerToken, " "); len(bearerTokenChars) == 2 {
		return bearerTokenChars[1]
	}

	return ""

}
