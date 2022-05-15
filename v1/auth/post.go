package auth

import (
	"net/http"
	"strconv"

	slwt "github.com/aswcloud/server-backend-local/jwt"
	"github.com/gin-gonic/gin"
)

func UserPost(c *gin.Context) {
	user := c.Params.ByName("user")
	_, err := strconv.Atoi(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "url 주소가 정상적인 요청이 아닙니다.",
		})
	}

	data, err := slwt.Create(map[string]interface{}{
		"id": user,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "정상적인 토큰 발행에 실패하였습니다.",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": data,
	})
}
