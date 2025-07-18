package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Role   string              `bson:"role" json:"role"`
	UserName string `bson:"username,omitempty" json:"username" validate:"required,min= 3,max= 50"`
	Email  string `bson:"email,omitempty" json:"email" validate:"required,min= 3,max= 50"`
	Password string `bson:"password,omitempty" json:"password" validate:"required,min= 6"`
	Phone string `bson:"phone,omitempty" json:"phone" validate:"required,min= 10"`
	Language string `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi Urdu English"`
	Age int `bson:"age,omitempty" json:"age" validate:"required,min= 10"`
	UserVerified struct {
		Email bool `bson:"emailVerified" json:"emailVerified"`
	} `bson:"userVerified" json:"userVerified"`
	UserVerifyToken struct {
		Email string `bson:"emailToken" json:"emailToken"`
		Phone string `bson:"phoneToken" json:"phoneToken"` 
	} `bson:"userVerifyToken" json:"userVerifyToken"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

}