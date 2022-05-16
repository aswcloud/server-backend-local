package admin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func uploadSingle(c *gin.Context, fileName string) (int, string) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error())
	}

	log.Println(file.Filename)

	// Upload the file to specific dst.
	// filename := filepath.Base(file.Filename)
	uploadPath := "./upload/" + fileName
	log.Println(uploadPath)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		return http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error())
	}

	return http.StatusOK, "success"
}

func UploadFile(c *gin.Context) {
	role, err := auth.Authorization(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if role.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "role id fail",
		})
		return
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	status, text := uploadSingle(c, uuid.String())
	c.JSON(status, gin.H{
		"msg": text,
	})
}
