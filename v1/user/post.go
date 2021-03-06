package user

import (
	"context"
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
	k8sjson "k8s.io/apimachinery/pkg/util/json"
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
	mdata, err := json.Marshal(data)
	log.Println(mdata)
	if err != nil {
		log.Println(err)
	}
	var deploy pbk8s.Deployment
	err = k8sjson.Unmarshal(mdata, &deploy)
	if err != nil {
		log.Println(err)
	}
	channel.CreateDeployment(context.TODO(), &deploy)
}
func createService(data map[string]interface{}, channel pbk8s.KubernetesClient) {
	mdata, err := json.Marshal(data)
	log.Println(mdata)
	if err != nil {
		log.Println(err)
	}
	var service pbk8s.Service
	err = k8sjson.Unmarshal(mdata, &service)
	if err != nil {
		log.Println(err)
	}
	channel.CreateService(context.TODO(), &service)
}

func CreatePersistentVolumeClaim(data map[string]interface{}, channel pbk8s.KubernetesClient) {
	mdata, err := json.Marshal(data)
	log.Println(string(mdata))
	if err != nil {
		log.Println(err)
	}
	var pvc pbk8s.Pvc
	err = k8sjson.Unmarshal(mdata, &pvc)
	if err != nil {
		log.Println(err)
	}
	channel.CreatePersistentVolumeClaim(context.TODO(), &pvc)
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
			createService(data, channel)
			break
		case "pvc":
			CreatePersistentVolumeClaim(data, channel)
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
			"msg": "????????? ????????? ????????????.",
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
