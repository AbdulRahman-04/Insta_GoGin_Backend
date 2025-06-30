package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId primitive.ObjectID `bson:"user_id,omitemtpy" json:"user_id"`
	Caption string `bson:"caption,omitempty" json:"caption" validate:"required,min= 3,max= 200"`
	ImageUrl string `bson:"image_url,omitempty" json:"image_url" validate:"required"`
	Tags   []string  `bson:"tags,omitempty" json:"tags" validate:"required,min= 2,max= 100"`
	Likes  int  `bson:"likes" json:"likes"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}