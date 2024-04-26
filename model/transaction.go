package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description"`
	Amount      float64            `json:"amount"`
}
