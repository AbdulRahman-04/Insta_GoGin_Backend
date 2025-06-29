// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // Post represents a private post created by a user
// type Post struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id" validate:"required"` // Linked user
// 	Caption   string             `bson:"caption" json:"caption" validate:"required,min=1,max=300"` // Optional max limit
// 	ImageURL  string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
// 	Tags      []string           `bson:"tags,omitempty" json:"tags,omitempty"` // Slice of tags
// 	Likes     int                `bson:"likes" json:"likes"`                   // Default 0
// 	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
// }

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID  `bson:"user_id,omitempty" json:"user_id"`
	ImageURL  string              `bson:"image_url,omitempty" json:"image_url" valdiate:"required"`
	Caption   string              `bson:"caption,omitempty" json:"caption,omitempty" validate:"required,min= 3, max= 300"`
	Tags      []string            `bson:"tags,omitempty" json:"tags" validate:"required,min= 2,max= 100"`
	Likes      int                `bson:"likes,omitempty" json:"likes" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`     
}