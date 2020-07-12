package routers

import "go-common/library/net/http/umr"

func NewRouter(engine *umr.Engine) {
	v1 := engine.Group("/api/v1")
	{
		v1.POST("/ping", nil)
	}
}
