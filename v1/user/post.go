package user

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	pbk8s "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-backend-local/database"
	ltemp "github.com/aswcloud/server-backend-local/template"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func getGrpcClient() (pbk8s.KubernetesClient, error) {
	k8s_server := os.Getenv("KUBERNETES_SERVER")
	conn, err := grpc.Dial(k8s_server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	log.Println(k8s_server)
	channel := pbk8s.NewKubernetesClient(conn)
	return channel, nil
}

func createDeployment(data map[string]interface{}, channel pbk8s.KubernetesClient) {

}
func createService(data map[string]interface{}, channel pbk8s.KubernetesClient) {

}

func CreatePersistentVolumeClaim(data map[string]interface{}, channel pbk8s.KubernetesClient) {

}

func jsonToK8s(jsonData string, id string) error {
	channel, err := getGrpcClient()
	if err != nil {
		return err
	}

	jsonData, err = ltemp.ConvertJsonToTemplate(jsonData, id)
	log.Println(jsonData)
	if err != nil {
		return err
	}
	var mapData []map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &mapData)
	if err != nil {
		return err
	}
	for _, data := range mapData {
		dataType := data["dataType"].(string)
		switch dataType {
		case "deployment":
			createDeployment(data, channel)
			break
		case "svc":
			createDeployment(data, channel)
			break
		case "pvc":
			createDeployment(data, channel)
			break
		}
	}
	log.Println(mapData)
	return err
}

func Post(c *gin.Context) {
	role, err := auth.Authorization(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if role.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "사용자 권한이 아닙니다.",
		})
		return
	}
	uuid := c.PostForm("uuid")

	db := database.New()
	db.Connect()
	defer db.Disconnect()

	dbTemp, err := db.Template().Get(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	log.Println(jsonToK8s(dbTemp.Json, role.Id))

	c.JSON(http.StatusOK, gin.H{
		"token": "token",
	})
}
