package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Admin represents admin collection schema in MongoDB
type Admin struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Name     string `bson:"name" json:"name" validate:"required,min=3,max=50"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password,omitempty" json:"-" validate:"required"`

	Phone string `bson:"phone" json:"phone" validate:"required"`


	AdminVerified struct {
		Email bool `bson:"email" json:"email"`
	} `bson:"adminVerified" json:"adminVerified"`

	AdminVerifyToken struct {
		Email string `bson:"email" json:"email"`
		Phone string `bson:"phone" json:"phone"`
	} `bson:"adminVerifyToken" json:"adminVerifyToken"`

	Role string `bson:"role" json:"role"` // Default: "admin" or "superadmin"

	CreatedAt int64 `bson:"createdAt" json:"createdAt"`
	UpdatedAt int64 `bson:"updatedAt" json:"updatedAt"`
}
