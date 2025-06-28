package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role string `bson:"role" json:"role"`
	AdminName string `bson:"adminName,omitempty" json:"adminName" validate:"required,min=3,max=50"`
	AdminEmail string `bson:"adminEmail,omitempty" json:"adminEmail" validate:"requored"`
	AdminPassword string `bson:"AdminPassword,omitempty" json:"adminPassword" validate:"required,min=6"`
	Phone string `bson:"phone,omitempty" json:"phone" validate:"required"`
	AdminVerified struct {
		Email bool `bson:"emailVerified" json:"emailVerified"`
	} `bson:"adminVerified" json:"adminVerified"`
	AdminVerifyToken struct {
		Email string `bson:"emailToken" json:"emailToken"`
		Phone string `bson:"phoneToken" json:"phoneToken"`
	}
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json: "updated_at"`
}