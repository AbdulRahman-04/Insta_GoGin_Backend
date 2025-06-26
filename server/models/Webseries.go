package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WebSeries struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SeriesName   string             `bson:"series_name,omitempty" json:"series_name" validate:"required"`
	Language     string             `bson:"language,omitempty" json:"language" validate:"required,oneof=Hindi English Telugu"`
	Genre        string             `bson:"genre,omitempty" json:"genre" validate:"required,oneof=Thriller Drama Romance Sci-Fi Action"`
	Seasons      int                `bson:"seasons,omitempty" json:"seasons" validate:"required,min=1"`
	Episodes     int                `bson:"episodes,omitempty" json:"episodes" validate:"required,min=1"`
	AverageTime  string             `bson:"average_time,omitempty" json:"average_time" validate:"required"`
	CreatedAt    int64              `bson:"created_at" json:"created_at"`
	UpdatedAt    int64              `bson:"updated_at" json:"updated_at"`
}