package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	pbk8s "github.com/aswcloud/idl/v1/kubernetes"
	slwt "github.com/aswcloud/server-backend-local/jwt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func createNamespace(id string) (bool, error) {
	k8s_server := os.Getenv("KUBERNETES_SERVER")
	conn, err := grpc.Dial(k8s_server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return false, err
	}

	log.Println(k8s_server)
	channel := pbk8s.NewKubernetesClient(conn)

	reply, err := channel.CreateNamespace(context.TODO(), &pbk8s.Namespace{
		Namespace: id,
	})
	if err != nil {
		return false, err
	}

	log.Println(reply, err)
	return true, err
}

func adminLogin(id, pw string) (int, gin.H) {
	if id == os.Getenv("ADMIN_ID") && pw == os.Getenv("ADMIN_PASSWORD") {
		data, err := slwt.Create(map[string]interface{}{
			"id":   id,
			"role": "admin",
		})
		if err != nil {
			return http.StatusBadRequest, gin.H{
				"msg": "정상적인 토큰 발행에 실패하였습니다.",
			}
		}

		return http.StatusOK, gin.H{
			"msg": data,
		}
	} else {
		return http.StatusBadRequest, gin.H{
			"msg": "관리자 계정 로그인에 실패하였습니다.",
		}
	}
}

func userLogin(id string) (int, gin.H) {
	_, err := strconv.Atoi(id)

	if err != nil {
		return http.StatusBadRequest, gin.H{
			"msg": "url 주소가 정상적인 요청이 아닙니다.",
		}
	}
	_, err = createNamespace(id)
	if err != nil {
		return http.StatusBadRequest, gin.H{
			"msg": "내부적인 서버 오류",
		}
	}

	data, err := slwt.Create(map[string]interface{}{
		"id":   id,
		"role": "user",
	})
	if err != nil {
		return http.StatusBadRequest, gin.H{
			"msg": "정상적인 토큰 발행에 실패하였습니다.",
		}
	}

	return http.StatusOK, gin.H{
		"msg": data,
	}
}

func UserPost(c *gin.Context) {
	user := c.Params.ByName("user")
	password := c.PostForm("password")

	if user != "" && password != "" {
		c.JSON(adminLogin(user, password))
	} else {
		c.JSON(userLogin(user))
	}
	return
}
