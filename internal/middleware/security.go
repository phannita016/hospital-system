package middleware

import "github.com/gin-gonic/gin"

func Security() gin.HandlerFunc {
	return func(g *gin.Context) {
		// X-XSS-Protection
		g.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

		// HTTP Strict Transport Security
		g.Writer.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		// X-Frame-Options
		g.Writer.Header().Add("X-Frame-Options", "SAMEORIGIN")

		// X-Content-Type-Options
		g.Writer.Header().Add("X-Content-Type-Options", "nosniff")

		// Content Security Policy
		g.Writer.Header().Add("Content-Security-Policy", "default-src 'self';")

		// X-Permitted-Cross-Domain-Policies
		g.Writer.Header().Add("X-Permitted-Cross-Domain-Policies", "none")

		// Referrer-Policy
		g.Writer.Header().Add("Referrer-Policy", "no-referrer")

		// Feature-Policy
		g.Writer.Header().Add("Feature-Policy", "microphone 'none'; camera 'none'")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(204)
			return
		}

		g.Next()
	}
}
