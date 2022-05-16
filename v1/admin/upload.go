package admin

import (
	"fmt"
	"log"

	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func uploadSingle(c *gin.Context, fileName string) error {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		return fmt.Errorf("get form err: %s", err.Error())
	}

	log.Println(file.Filename)

	// Upload the file to specific dst.
	// filename := filepath.Base(file.Filename)
	uploadPath := "./upload/" + fileName
	log.Println(uploadPath)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		return fmt.Errorf("upload file err: %s", err.Error())
	}

	return nil
}

func UploadFile(c *gin.Context) (string, error) {
	role, err := auth.Authorization(c)
	if err != nil {
		return "", err
	}
	if role.Role != "admin" {
		return "", fmt.Errorf("권한이 부족합니다.")
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	if err := uploadSingle(c, uuid.String()); err != nil {
		return "", err
	}

	return uuid.String(), nil
}
