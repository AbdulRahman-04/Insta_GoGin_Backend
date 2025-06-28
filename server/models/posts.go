package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"` // Kis user ka post hai
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	ImageURL  string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	Tags      []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
