package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Template struct {
	collection      *mongo.Collection
	totalCollection *mongo.Collection
}

func (self *Template) Add() {

}

func (self *Template) Delete() {

}

func (self *Template) Update() {

}
