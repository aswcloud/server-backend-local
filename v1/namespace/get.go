package namespace

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	user := c.Params.ByName("user")

	c.JSON(http.StatusOK, gin.H{
		"msg": user,
	})
}
