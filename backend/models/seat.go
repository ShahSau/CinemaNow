package models

import "time"

type Seat struct {
	Id           int       `gorm:"primary_key" json:"id"`
	AuditoriumId int       `gorm:"column:auditorium_id" json:"auditorium_id"`
	Row          int       `gorm:"column:row" json:"row"`
	Number       int       `gorm:"column:number" json:"number"`
	Available    bool      `gorm:"column:available" json:"available"`
	Price        int64     `gorm:"column:price" json:"price"`
	Type         string    `gorm:"column:type" json:"type"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
