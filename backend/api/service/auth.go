package service

import (
	"kursarbeit/api/my_jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.IndentedJSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err := my_jwt.ValidateToken(tokenString)
		if err != nil {
			ctx.IndentedJSON(401, gin.H{"error": "Unable to validate token: " + err.Error()})
			ctx.Abort()
			return
		}
		// id, err := my_jwt.ExtractID(tokenString)

		ctx.Next()
	}
}
