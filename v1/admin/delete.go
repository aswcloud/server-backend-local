package admin

import (
	"net/http"

	"github.com/aswcloud/server-backend-local/database"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
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

	uuid := c.PostForm("uuid")

	db := database.New()
	db.Connect()
	defer db.Disconnect()

	err = db.Template().Delete(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "데이터가 옳바르지 않은 데이터를 삭제하려고 하였습니다.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "성공적으로 삭제 하였습니다.",
		})
	}
}
