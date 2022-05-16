package user

import (
	"net/http"
	"os"

	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetImage(c *gin.Context) {
	_, err := auth.Authorization(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	uuid := c.Params.ByName("uuid")

	if fileExists("./upload/" + uuid) {
		c.Status(200)
		c.File("./upload/" + uuid)
	} else {
		c.Status(400)
		c.File(("./resources/empty.png"))
	}
}
