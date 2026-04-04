package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_header := ctx.GetHeader("Authorization")
		token := strings.Split(token_header, " ")[1]
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		session, err := m.service.GetSession(ctx.Request.Context(), token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		ctx.Set(UserIDKey, session.Edges.Owner.ID)
		ctx.Next()
	}
}
