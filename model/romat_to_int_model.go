package model

type RomanToIntRequest struct {
	A string `validate:"required" json:"a" example:"III"`
}
