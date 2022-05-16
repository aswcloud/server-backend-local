package admin

import (
	"net/http"

	"github.com/aswcloud/server-backend-local/database"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
)

func Put(c *gin.Context) {
	role, err := auth.Authorization(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if role.Role != "admin" {
		c.JSON(http.StatusBadRequest, "권한이 부족합니다.")
		return
	}

	name := c.PostForm("name")
	jsonData := c.PostForm("json")
	uuid, err := UploadFile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	db := database.New()
	db.Connect()
	defer db.Disconnect()

	db.Template().Update(uuid, name, jsonData)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
