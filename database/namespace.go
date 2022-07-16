package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Namespace struct {
	collection *mongo.Collection
}

func (self *Client) Namespace() *Namespace {
	name := Namespace{}
	if self.database == nil {
		return nil
	}
	name.collection = self.database.Collection("namespace")
	return &name
}
