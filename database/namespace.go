package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Namespace struct {
	collection      *mongo.Collection
	totalCollection *mongo.Collection
}
