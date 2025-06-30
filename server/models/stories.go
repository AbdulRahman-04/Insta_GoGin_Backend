package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Stories struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Caption   string             `bson:"caption,omitempty" json:"caption" validate:"required"`
	ImageUrl  string             `bson:"imageUrl,omitempty" json:"imageUrl" validate:"required"`
	Text      string             `bson:"text,omitempty" json:"text" validate:"required"`
	Song      string             `bson:"song,omitempty" json:"song" validate:"required,oneof=Hindi English BrazilianFonk"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
