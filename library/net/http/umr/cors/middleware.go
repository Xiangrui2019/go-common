package cors

import (
	"go-common/library/net/http/umr"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(cors_domain string) umr.HandlerFunc {
	config := DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}

	if gin.Mode() != gin.ReleaseMode {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = []string{cors_domain}
	}
	config.AllowCredentials = true

	return New(config)
}
