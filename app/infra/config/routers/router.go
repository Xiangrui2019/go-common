package routers

import (
	"go-common/app/infra/config/api"
	"go-common/library/net/http/umr"
)

func NewRouter(engine *umr.Engine) {
	v1 := engine.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		admin := v1.Group("/admin")
		{
		}

		configing := v1.Group("/configing")
		{
		}
	}
}
