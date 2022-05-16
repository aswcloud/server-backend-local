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
		vauth := v1.Group("/auth")
		{
			vauth.POST("/:user", v1auth.UserPost)
		}
		gadmin := v1.Group("/admin")
		{
			gadmin.GET("/:user", admin.Get)
			gadmin.POST("/:user", admin.Post)
			gadmin.DELETE("/:user", admin.Delete)
			gadmin.POST("/upload", admin.UploadFile)
		}
		guser := v1.Group("/user")
		{
			guser.GET("/:user", user.Get)
			guser.POST("/:user", user.Post)
			guser.DELETE("/:user", user.Delete)
			guser.GET("/upload/:uuid", user.GetImage)
		}
	}

	r.Run(":8080")
}
