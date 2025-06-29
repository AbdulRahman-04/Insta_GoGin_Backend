package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Story struct {
	ID            primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID    `bson:"user_id,omitempty" json:"user_id"`
	ImageURL      string                `bson:"image_url,omitempty" json:"image_url"`
	Text          string                `bson:"text,omitempty" json:"text" validate:"max= 100"`
	Likes         int                   `bson:"likes" json:"likes"`
	CreatedAt     time.Time             `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time             `bson:"updated_at" json:"updated_at"`    
}