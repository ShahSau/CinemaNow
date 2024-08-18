package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie is a struct that represents the movie model
type Ticket struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Date      string             `json:"date" bson:"date" validate:"required"`
	Time      string             `json:"time" bson:"time" validate:"required"`
	Day       string             `json:"day" bson:"day" validate:"required"`
	Price     string             `json:"price" bson:"price" validate:"required"`
	Location  string             `json:"location" bson:"location" validate:"required"`
	Seats     string             `json:"seats" bson:"seats" validate:"required"`
	Row       string             `json:"row" bson:"row" validate:"required"`
	Theatre   string             `json:"theatre" bson:"theatre" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
