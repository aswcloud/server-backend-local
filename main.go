package main

import (
	"log"

	"github.com/aswcloud/server-backend-local/database"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	database := database.New()
	if !database.Connect() {
		log.Println("database doesnt connect")
		log.Panic()
	}
	database.Disconnect()

	r := gin.Default()

	r.Group("/v1")
	{
		ns := r.Group("namespace")
		{
			ns.GET("/:user", func(ctx *gin.Context) {

			})
		}
	}

	r.Run(":8080")
}
