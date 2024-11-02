package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Tags      []string           `bson:"tags" json:"tags"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
