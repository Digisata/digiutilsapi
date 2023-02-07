package model

type ContainsDuplicateRequest struct {
	A []int `validate:"required" json:"a" example:"1,2,3,1"`
}
