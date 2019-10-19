package models

// Gender type
type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

// People holds personal information
type People struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    Gender `json:"gender"`
}
