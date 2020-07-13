package routers

import (
	"go-common/app/infra/config/api"
	"go-common/library/net/http/umr"
	"go-common/library/net/http/umr/cors"
	"os"
)

func NewRouter(engine *umr.Engine) {
	v1 := engine.Group("/api/v1")
	{
		v1.Use(cors.CorsMiddleware(os.Getenv("CORS_DOMAIN")))

		v1.POST("/ping", api.Ping)

		admin := v1.Group("/admin")
		{
		}

		configing := v1.Group("/configing")
		{
		}
	}
}
