package types

// Auditorium is a struct that represents the auditorium model
type Auditorium struct {
	ID            string `json:"_id,omitempty"`
	Name          string `json:"name" validate:"required"`
	MovieID       string `json:"movie_id" validate:"required"`
	NoSeats       int    `json:"no_seats" validate:"required"`
	Rows          int    `json:"rows" validate:"required"`
	Columns       int    `json:"columns" validate:"required"`
	Seats         []Seat `json:"seats" bson:"seats" validate:"required"`
	SelectedSeats []Seat `json:"selected_seats" bson:"selected_seats" validate:"required"`
}
