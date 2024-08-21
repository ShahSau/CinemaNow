package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TransactionId int                `json:"transaction_id"`
	UserId        int                `json:"user_id"`
	SeatId        int                `json:"seat_id"`
	ScreeningId   int                `json:"screening_id"`
	Status        int                `json:"status"`
	Transaction   Transaction        `json:"transaction,omitempty"`
	Seat          Seat               `json:"seat"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
