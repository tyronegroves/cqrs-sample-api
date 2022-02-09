package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const UsernameKey = "Username"

func (s *server) username(ctx *gin.Context) string {
	return ctx.GetString(UsernameKey)
}

func (s *server) authorizeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer"
		auth := ctx.GetHeader("Authorization")

		if len(auth) < len(BearerSchema)+4 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(UsernameKey, strings.TrimSpace(auth[len(BearerSchema):]))
	}
}
