package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MovieName  string             `bson:"movie_name,omitempty" json:"movie_name" validate:"required"`
	Language   string             `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Telugu"`
	Duration   string             `bson:"duration,omitempty" json:"duration" validate:"required"`
	MovieGenre string             `bson:"movie_genre,omitempty" json:"movie_genre" validate:"required,oneof=Horror Comedy Drama Action"`
	CreatedAt  int64              `bson:"created_at" json:"created_at"`
	UpdatedAt  int64              `bson:"updated_at" json:"updated_at"`
}