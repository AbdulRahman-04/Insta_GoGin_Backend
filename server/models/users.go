package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role string `bson:"role" json:"role"`
	UserName string `bson:"userName,omitempty" json:"userName" validate:"required,min= 3,max= 50"`
	Email string `bson:"email,omitempty" json:"email" validate:"required"`
	Password string `bson:"password,omitempty" json:"password" validate:"required,min= 6"`
    Phone string `bson:"phone,omitempty" json:"phone" validate:"required,min= 10"`
	Age int `bson:"age,omitempty" json:"age" validate:"required,min= 10"`
	Language string `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Urdu Kannada Arabic"`
	UserVerified struct {
		Email bool `bson:"emailVerified" json:"emailVerified"`
	} `bson:"userVerified" json:"userVerified"`
   UserVerifyToken struct {
	Email string `bson:"email" json:"email"`
	Phone string `bson:"phone" json:"phone"`
   } `bson:"userVerifyToken" json:"userVerifyToken"`
   CreatedAt time.Time `bson:"created_at" json:"created_at"`
   UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}