package types

// Booking is a struct that represents the booking model
type Booking struct {
	ID            string `json:"_id,omitempty"`
	TransactionId int    `json:"transaction_id"`
	UserId        int    `json:"user_id"`
	SeatId        int    `json:"seat_id"`
	ScreeningId   int    `json:"screening_id"`
	Status        bool   `json:"status"`
	Seat          []Seat `json:"seat"`
}
