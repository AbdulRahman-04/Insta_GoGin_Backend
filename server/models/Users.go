package models

import (
	  "go.mongodb.org/mongo-driver/bson/primitive"
       "time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role string `bson:"role" json:"role"`
	UserName string `bson:"username,omitempty" json:"username" validate:"required,max=50,min=3"`
	Email string `bson:"email,omitempty" json:"email" validate:"required"`
	Password string `bson:"password,omitempty" json:"password" validate:"required"`
	Phone string `bson:"phone,omitempty" json:"phone" validate:"required"`
	Age int `bson:"age" json:"age"`
	Language string `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Urdu Kannada Tamil"`
	UserVerified struct {
		Email string `bson:"email" json:"email"`
	} `bson:"userverified" json:"userverified"`
	UservErifyToken struct {
		Email string `bson:"emailtoken" json:"emailtoken"`
		Phone string `bson:"phonetoken" json:"phonetoken"`
	} `bson:"userVerifyToken" json:"userVerifyToken"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}