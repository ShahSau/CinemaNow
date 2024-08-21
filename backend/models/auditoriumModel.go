package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auditorium is a struct that represents the auditorium model
type Auditorium struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" validate:"required"`
	MovieID string             `json:"movie_id" validate:"required"`
	NoSeats int                `json:"no_seats" validate:"required"`
	Rows    int                `json:"rows" validate:"required"`
	Columns int                `json:"columns" validate:"required"`

	Seats         []Seat `json:"seats" bson:"seats" validate:"required"`
	SelectedSeats []Seat `json:"selected_seats" bson:"selected_seats" validate:"required"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
