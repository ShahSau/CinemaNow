package types

// Theater is a struct that represents the theater model
type Theater struct {
	ID         int          `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string       `json:"name" validate:"required"`
	Address    string       `json:"address" validate:"required"`
	Auditorium []Auditorium `json:"auditorium"`
}
