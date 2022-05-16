package admin

import (
	"net/http"

	"github.com/aswcloud/server-backend-local/database"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_, err := auth.Authorization(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	db := database.New()
	db.Connect()
	defer db.Disconnect()

	data := db.Template().List()
	c.JSON(http.StatusOK, gin.H{
		"msg": data,
	})
}
