package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})
			ctx.Abort()
			return
		}

		token := parts[1]

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token missing",
			})
			ctx.Abort()
			return
		}

		session, err := m.service.GetSession(ctx.Request.Context(), token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		ctx.Set(UserIDKey, session.Edges.Owner.ID)
		ctx.Next()
	}
}
