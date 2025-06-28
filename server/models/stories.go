package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Story struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	MediaURL  string             `bson:"media_url" json:"media_url"`
	Caption   string             `bson:"caption,omitempty" json:"caption,omitempty"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"` // like 24 hours
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
