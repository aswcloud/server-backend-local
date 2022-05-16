package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"github.com/aswcloud/server-backend-local/v1/admin"
	v1auth "github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/aswcloud/server-backend-local/v1/user"
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
		guser := v1.Group("/user")
		{
			guser.GET("/:user", user.Get)
			guser.POST("/:user", user.Post)
			guser.DELETE("/:user", user.Delete)
		}
		gadmin := v1.Group("/admin")
		{
			gadmin.GET("/:user", admin.Get)
			gadmin.POST("/:user", admin.Post)
			gadmin.DELETE("/:user", admin.Delete)
		}
	}

	r.Run(":8080")
}
