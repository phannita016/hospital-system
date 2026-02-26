package middleware

import (
	"errors"
	"hospital/internal/modules/staff/service"
	"hospital/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("Authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("Invalid authorization format"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := tokenService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			utils.ErrorResponse(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*service.Claims)
		if !ok {
			utils.ErrorResponse(c, http.StatusUnauthorized, errors.New("Invalid token claims"))
			c.Abort()
			return
		}

		c.Set("staff_id", claims.StaffID)
		c.Set("hospital_id", claims.HospitalID)
		c.Next()
	}
}
