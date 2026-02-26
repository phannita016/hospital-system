package middleware

import (
	"hospital/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 || c.Writer.Status() >= 400 {
			return
		}

		resp := c.Keys["response"]
		if resp != nil {
			c.JSON(http.StatusOK, utils.APIResponse{
				Code:    http.StatusOK,
				Message: "SUCCESS",
				Data:    resp,
			})
		}
	}
}
