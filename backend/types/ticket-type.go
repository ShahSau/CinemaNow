package types

// Ticket is a struct that represents the ticket model
type Ticket struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Date     string `json:"date" bson:"date" validate:"required"`
	Time     string `json:"time" bson:"time" validate:"required"`
	Day      string `json:"day" bson:"day" validate:"required"`
	Price    string `json:"price" bson:"price" validate:"required"`
	Location string `json:"location" bson:"location" validate:"required"`
	Seats    string `json:"seats" bson:"seats" validate:"required"`
	Row      string `json:"row" bson:"row" validate:"required"`
	Theatre  string `json:"theatre" bson:"theatre" validate:"required"`
}
