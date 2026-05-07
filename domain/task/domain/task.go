package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	ID        bson.ObjectID `bson:"_id"        json:"id"`
	Title     string        `bson:"title"      json:"title"`
	Status    Status        `bson:"status"     json:"status"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

type Status string
