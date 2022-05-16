package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbTemplate struct {
	Uuid     string
	Name     string
	Json     string
	CreateAt string
	ImageId  string
}

type Template struct {
	collection *mongo.Collection
}

func (self *Client) Template() *Template {
	template := Template{}
	if self.database == nil {
		return nil
	}
	template.collection = self.database.Collection("template")
	return &template
}

func (self *Template) Add(name, data, imageId string) error {
	t := time.Now().Format(time.RFC3339)
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	self.collection.InsertOne(context.TODO(), bson.D{
		{"uuid", u.String()},
		{"name", name},
		{"json", data},
		{"createAt", t},
		{"imageId", imageId},
	})
	return nil
}

func (self *Template) Delete(uuid string) error {
	// TODO:
	// uuid를 통해서 Collection을 탐색한 뒤, imageId 값을 구한 뒤 삭제 해야함

	_, err := self.collection.DeleteMany(context.TODO(), bson.D{
		{"uuid", uuid},
	})
	return err
}

func (self *Template) Update(uuid, name, json string) {
	self.collection.UpdateOne(context.TODO(), bson.D{
		{"uuid", uuid},
	}, bson.D{
		{"name", name},
		{"json", json},
		// {"imageId", imageId},
	})
}

func (self *Template) List() []DbTemplate {
	cursor, err := self.collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Println(err)
		return []DbTemplate{}
	}
	db := []DbTemplate{}
	for cursor.Next(context.TODO()) {
		var elem bson.M
		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}
		// find 결과 print
		fmt.Println(elem)
		db = append(db, DbTemplate{
			Uuid:     elem["uuid"].(string),
			Name:     elem["name"].(string),
			Json:     elem["json"].(string),
			CreateAt: elem["createAt"].(string),
			ImageId:  elem["imageId"].(string),
		})
	}

	return db
}
