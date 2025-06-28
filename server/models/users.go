package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Users struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role string `bson:"role" json:"role"`
	UserName string `bson:"username,omitempty" json:"username" validate:"required,min=3,max=50"`
	Email string `bson:"email,omitempty" json:"email" validate:"required"`
	Password string `bson:"password,omitempty" json:"password"`
	Language string `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Kannada"`
	UserVerified struct {
		Email bool `bson:"email" json:"email"`
	} `bson:"userverified" json:"userverified"`
	UserVerifyToken struct {
		Email string `bson:"email" json:"email"`
		Phone string `bson:"phone" json:"phone"`
	} `bson:"userverifytoken" json:"userverifytoken"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}