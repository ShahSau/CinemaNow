package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Screening struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AuditoriumId int                `json:"auditorium_id" validate:"required"`
	CinemaId     int                `json:"cinema_id" validate:"required"`
	MovieId      int                `json:"movie_id" validate:"required"`
	StartTime    time.Time          `json:"start_time" validate:"required"`
	Auditorium   []Auditorium       `json:"auditorium"`
	Theater      []Theater          `json:"theater"`
	Movie        []Movie            `json:"movie"`
	Bookable     bool               `json:"bookable"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
