package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

var CollectionName string = "learningMongo"

type UserCredentials struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
