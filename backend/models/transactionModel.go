package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction is a struct that represents the transaction model
type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TicketID  int                `json:"ticket_id" bson:"ticket_id" validate:"required"`
	UserID    int                `json:"user_id" bson:"user_id" validate:"required"`
	Quantity  int                `json:"quantity" bson:"quantity" validate:"required"`
	Total     float64            `json:"total" bson:"total" validate:"required"`
	Paid      bool               `json:"paid" bson:"paid" validate:"required"`
	Ticket    []Ticket           `json:"ticket,omitempty" bson:"ticket,omitempty"`
	Booking   []Booking          `json:"booking,omitempty" bson:"booking,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
