package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AdminName string `bson:"adminName,omitempty" json:"adminName" validate:"required,min=3 ,max=50"`
	AdminEmail string `bson:"adminEmail,omitempty" json:"email" validate:"required"`
	Password string `bson:"password,omitempty" json:"password" validate:"required,min=6"`
	Language string `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Urdu Tamil"`
	Phone string `bson:"phone,omitempty" json:"phone" validate:"required,min= 10"`
	AdminVerified struct {
		Email string `bson:"email" json:"email"`
	} `bson:"adminVerified" json:"adminVerified"`
	AdminVerifyToken struct {
		Email string `bson:"emailToken" json:"emailToken"`
		Phone string `bson:"phoneString" json:"phoneString"`
	}
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}