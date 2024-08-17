package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie is a struct that represents the movie model
type Movie struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type             string             `json:"type" bson:"type" validate:"required"`
	Adult            bool               `json:"adult" validate:"required"`
	BackdropPath     string             `json:"backdrop_path" validate:"required"`
	GenreIds         []int              `json:"genre_ids" validate:"required"`
	OriginalLanguage string             `json:"original_language" validate:"required"`
	OriginalTitle    string             `json:"original_title" validate:"required"`
	Overview         string             `json:"overview" validate:"required"`
	Popularity       float64            `json:"popularity" validate:"required"`
	PosterPath       string             `json:"poster_path" validate:"required"`
	ReleaseDate      string             `json:"release_date" validate:"required"`
	Title            string             `json:"title" validate:"required"`
	Video            bool               `json:"video" validate:"required"`
	VoteAverage      float64            `json:"vote_average" validate:"required"`
	VoteCount        int                `json:"vote_count" validate:"required"`
	MovieID          int                `json:"movie_id,omitempty" bson:"movie_id" validate:"required" unique:"true"`
	CreatedAt        time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt        time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
