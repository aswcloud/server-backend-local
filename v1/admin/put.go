package admin

import (
	"net/http"
	"strings"

	"github.com/aswcloud/server-backend-local/jwt"
	"github.com/gin-gonic/gin"
)

func Put(c *gin.Context) {
	bearer := strings.Split(c.GetHeader("Authorization"), " ")
	// c.PostForm("phone")

	if len(bearer) != 2 || bearer[0] != "Bearer" {
		// return "", fmt.Errorf("Authorization: bearer not match")
	}
	token, err := jwt.Validate(bearer[1])
	if err != nil {
		// return "", err
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
