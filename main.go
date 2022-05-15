package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	v1auth "github.com/aswcloud/server-backend-local/v1/auth"
	v1deploy "github.com/aswcloud/server-backend-local/v1/deployment"
	v1ns "github.com/aswcloud/server-backend-local/v1/namespace"
	v1svc "github.com/aswcloud/server-backend-local/v1/service"
	v1storage "github.com/aswcloud/server-backend-local/v1/storage"
)

func main() {
	gotenv.Load()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/:user", v1auth.UserPost)
		}
		ns := v1.Group("/namespace")
		{
			ns.GET("/:user", v1ns.Get)
			ns.POST("/:user", v1ns.Post)
			ns.DELETE("/:user", v1ns.Delete)
		}
		deploy := v1.Group("/deployment")
		{
			deploy.GET("/:user", v1deploy.Get)
			deploy.POST("/:user", v1deploy.Post)
			deploy.DELETE("/:user", v1deploy.Delete)
		}
		svc := v1.Group("/service")
		{
			svc.GET("/:user", v1svc.Get)
			svc.POST("/:user", v1svc.Post)
			svc.DELETE("/:user", v1svc.Delete)
		}
		storage := v1.Group("/storage")
		{
			storage.GET("/:user", v1storage.Get)
			storage.POST("/:user", v1storage.Post)
			storage.DELETE("/:user", v1storage.Delete)
		}
	}

	r.Run(":8080")
}
