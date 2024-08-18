package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MovieDetails is a struct that represents the movie details model
type MovieDetails struct {
	ID                  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Adult               bool               `json:"adult" validate:"required"`
	BackdropPath        string             `json:"backdrop_path" validate:"required"`
	BelongsToCollection interface{}        `json:"belongs_to_collection" bson:"belongs_to_collection"`
	Budget              int                `json:"budget" validate:"required"`
	Genres              []struct {
		ID   int    `json:"id" validate:"required"`
		Name string `json:"name" validate:"required"`
	} `json:"genres" validate:"required"`
	Homepage            string   `json:"homepage" validate:"required"`
	ImdbID              string   `json:"imdb_id" validate:"required"`
	MovieID             int      `json:"movie_id" validate:"required" unique:"true"`
	OriginalCountry     []string `json:"original_country" validate:"required"`
	OriginalLanguage    string   `json:"original_language" validate:"required"`
	OriginalTitle       string   `json:"original_title" validate:"required"`
	Overview            string   `json:"overview" validate:"required"`
	Popularity          float64  `json:"popularity" validate:"required"`
	PosterPath          string   `json:"poster_path" validate:"required"`
	ProductionCompanies []struct {
		ID            int    `json:"id" validate:"required"`
		LogoPath      string `json:"logo_path" validate:"required"`
		Name          string `json:"name" validate:"required"`
		OriginCountry string `json:"origin_country" validate:"required"`
	} `json:"production_companies" validate:"required"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1" validate:"required"`
		Name     string `json:"name" validate:"required"`
	} `json:"production_countries" validate:"required"`
	ReleaseDate     string `json:"release_date" validate:"required"`
	Revenue         int    `json:"revenue" validate:"required"`
	Runtime         int    `json:"runtime" validate:"required"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name" validate:"required"`
		Iso6391     string `json:"iso_639_1" validate:"required"`
		Name        string `json:"name" validate:"required"`
	} `json:"spoken_languages" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Tagline     string    `json:"tagline" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Video       bool      `json:"video" validate:"required"`
	VoteAverage float64   `json:"vote_average" validate:"required"`
	VoteCount   int       `json:"vote_count" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
