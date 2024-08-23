package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Screening is a struct that represents the Screening model
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
}
