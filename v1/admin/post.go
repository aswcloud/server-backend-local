package admin

import (
	"net/http"

	"github.com/aswcloud/server-backend-local/database"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
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
	// base64 인코딩 됨.
	jsonData := c.PostForm("json")
	// rawJson, err := base64.StdEncoding.DecodeString(jsonData)

	uuid, err := UploadFile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	db := database.New()
	db.Connect()
	defer db.Disconnect()

	db.Template().Add(name, jsonData, uuid)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
