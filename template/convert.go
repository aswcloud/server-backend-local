package template

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	temp "text/template"

	"github.com/gofrs/uuid"
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

func ConvertJsonToTemplate(jsonData string, id string) (string, error) {
	dbjson := k8sTemplate{}
	err := json.Unmarshal([]byte(jsonData), &dbjson)
	log.Println(dbjson)
	dbjson.Namespace = id
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	dbjson.DeploymentName += ("-" + id + "-" + uuid.String())
	dbjson.TemplateName += ("-" + id + "-" + uuid.String())
	dbjson.PvcName += ("-" + id + "-" + uuid.String())

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
