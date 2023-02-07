package model

type AddBinaryRequest struct {
	A string `validate:"required" json:"a" example:"1101"`
	B string `validate:"required" json:"b" example:"11"`
}
