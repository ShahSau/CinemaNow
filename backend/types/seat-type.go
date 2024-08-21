package types

// Seat is a struct that represents the seat model
type Seat struct {
	ID           int    `json:"id"`
	AuditoriumId int    `json:"auditorium_id"`
	Row          int    `json:"row"`
	Number       int    `json:"number"`
	Available    bool   `json:"available"`
	Price        int64  `json:"price"`
	Type         string `json:"type"`
}
