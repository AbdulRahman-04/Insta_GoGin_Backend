package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role     string             `bson:"role" json:"role"` // Should always be "admin"
	AdminName     string             `bson:"adminname,omitempty" json:"adminname" validate:"required,min=3,max=50"`
	Email    string             `bson:"email,omitempty" json:"email" validate:"required,email"`
	Password string             `bson:"password,omitempty" json:"password" validate:"required,min=6"`
	Phone    string             `bson:"phone,omitempty" json:"phone" validate:"required,min=10,max=15"`
	Language string             `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi Urdu English"`
	Age      int                `bson:"age,omitempty" json:"age" validate:"required,min=18,max=100"`

	AdminVerified struct {
		Email bool `bson:"emailVerified" json:"emailVerified"`
	} `bson:"adminVerified" json:"adminVerified"`
	AdminVerfiyToken struct {
		Email string `bson:"emailToken" json:"emailToken"`
		Phone string `bson:"phoneToken" json:"phoneToken"` 
	} `bson:"adminVerifyToken" json:"adminVerifyToken"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
