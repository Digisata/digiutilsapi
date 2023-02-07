package model

type IsPowerOfTwoRequest struct {
	A int `validate:"required" json:"a" example:"12"`
}