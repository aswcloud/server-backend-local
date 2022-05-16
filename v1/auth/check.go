package auth

import (
	"fmt"
	"strings"

	"github.com/aswcloud/server-backend-local/jwt"
	"github.com/gin-gonic/gin"
)

type Role struct {
	Id   string
	Role string
}

func Authorization(c *gin.Context) (Role, error) {
	bearer := strings.Split(c.GetHeader("Authorization"), " ")
	if len(bearer) != 2 || bearer[0] != "Bearer" {
		return Role{}, fmt.Errorf("Authorization is not found")
	}
	token, err := jwt.Validate(bearer[1])
	if err != nil {
		return Role{}, fmt.Errorf("Access token expired")
	}

	return Role{
		Id:   token["id"].(string),
		Role: token["role"].(string),
	}, nil
}
