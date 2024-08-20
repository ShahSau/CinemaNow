package models

import (
	"time"
)

// Auditorium is a struct that represents the auditorium model
type Auditorium struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	MovieID string `gorm:"column:cinema_id" json:"movie_id"`
	NoSeats int    `gorm:"column:no_seats" json:"no_seats"`
	Rows    int    `gorm:"column:rows" json:"rows"`
	Columns int    `gorm:"column:columns" json:"columns"`

	Seats         []Seat `gorm:"foreignKey:AuditoriumId" json:"seats"`
	SelectedSeats []Seat `json:"selected_seats"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
