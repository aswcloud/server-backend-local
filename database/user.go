package database

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type UserInfo struct {
	Nickname string
	Email    string
	UserId   string
	Uuid     string
}

type UserCollection struct {
	collection *mongo.Collection
}

func (self *Client) UserCollection() *UserCollection {
	user := UserCollection{}
	if self.database == nil {
		return nil
	}

	user.collection = self.database.Collection("register")
	return &user
}

func (self UserCollection) ExistsId(userId string) bool {
	result := self.collection.FindOne(context.TODO(), bson.D{
		{"userId", userId},
	})
	var elem bson.D
	err := result.Decode(&elem)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (self UserCollection) Register(userId, password, phone string) (string, error) {
	if self.ExistsId(userId) {
		return "", fmt.Errorf("exists userid : " + userId)
	}

	hash := sha512.Sum512([]byte(password))
	text := hex.EncodeToString(hash[:])
	b_uuid, _ := uuid.New()
	uuid := hex.EncodeToString(b_uuid[:])
	timeStamp := time.Now().Format(time.RFC3339)
	self.collection.InsertOne(context.TODO(), bson.D{
		{"uuid", uuid},
		{"userId", userId},
		{"password", text},
		{"created_at", timeStamp},
		{"phone", phone},
	})

	return uuid, nil
}

func (self UserCollection) Login(userId, password string) bool {
	hash := sha512.Sum512([]byte(password))
	text := hex.EncodeToString(hash[:])

	result := self.collection.FindOne(context.TODO(), bson.D{
		{"userId", userId},
		{"password", text},
	})
	var elem bson.D
	err := result.Decode(&elem)
	if err != nil {
		return false
	} else {
		return true
	}
}
