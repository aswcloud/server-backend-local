package user

import (
	"context"
	"log"
	"net/http"
	"os"

	pbk8s "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-backend-local/database"
	"github.com/aswcloud/server-backend-local/v1/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
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

func jsonToK8s(jsonData string) error {
	channel, err := getGrpcClient()
	if err != nil {
		return err
	}

	data := pbk8s.Deployment{
		Namespace:    "20183221",
		Name:         "chacha-deployment",
		TemplateName: "chacha-app",
		ReplicaCount: 1,
		Volume: []*pbk8s.DeploymentVolume{
			{
				Name:      "pvc-chacha",
				ClaimName: "pvc-chacha",
			},
		},
		Containers: []*pbk8s.DeploymentContainer{
			{
				Name:  "chahca-app",
				Image: "nginx:1.14.2",
				Ports: []int32{80},
				VolumeMount: []*pbk8s.DeploymentVolumemount{
					{
						Name:      "pvc-chacha",
						MountPath: "/usr/share/nginx/html",
					},
				},
			},
		},
	}
	pvc := pbk8s.Pvc{
		Namespace:        "20183221",
		Name:             "pvc-chacha",
		Capacity:         "2Gi",
		StorageClassName: "hostpath",
		AccessMode:       []string{"ReadWriteOnce"},
	}

	bytee, _ := protojson.Marshal(&data)
	log.Println(string(bytee))

	bytee, _ = protojson.Marshal(&pvc)
	log.Println(string(bytee))

	result, err := channel.CreatePersistentVolumeClaim(context.TODO(), &pvc)
	if err != nil {
		log.Println("err : ", err)
	}
	log.Println(result)
	result, err = channel.CreateDeployment(context.TODO(), &data)
	if err != nil {
		log.Println("err : ", err)
	}
	log.Println(result)
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

	_, err = db.Template().Get(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Println(jsonToK8s(""))

	c.JSON(http.StatusOK, gin.H{
		"token": "token",
	})
}
