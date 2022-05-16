package template

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	temp "text/template"
)

type k8sTemplate struct {
	Image            string `json:"image"`
	Namespace        string `json:"namespace"`
	DeploymentName   string `json:"deploymentName"`
	TemplateName     string `json:"templateName"`
	PvcName          string `json:"pvcName"`
	PvcPath          string `json:"pvcPath"`
	Capacity         string `json:"capacity"`
	StorageClassName string `json:"storageClass"`
	ContainerPort    int    `json:"containerPort"`
}

func Uuid(len int) string {
	result := ""
	for i := 0; i < len; i++ {
		result += strconv.Itoa(int(rand.Int31n(10)))
	}
	return result
}

func ConvertJsonToTemplate(jsonData string, id string) (string, error) {
	dbjson := k8sTemplate{}
	err := json.Unmarshal([]byte(jsonData), &dbjson)
	log.Println(dbjson)
	dbjson.Namespace = id
	uuid := Uuid(8)
	if err != nil {
		return "", err
	}

	dbjson.DeploymentName += ("-" + id + "-" + uuid)
	dbjson.TemplateName += ("-" + id + "-" + uuid)
	dbjson.PvcName += ("-" + id + "-" + uuid)

	a := temp.New("test")
	data, err := os.ReadFile("./template/data.json")
	if err != nil {
		return "", err
	}
	a, err = a.Parse(string(data))
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	a.Execute(&sb, dbjson)

	return sb.String(), nil
}
