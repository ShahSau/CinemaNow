package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Theater is a struct that represents the theater model
type Theater struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" validate:"required"`
	Address    string             `json:"address" validate:"required"`
	Auditorium []Auditorium       `json:"auditorium"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
