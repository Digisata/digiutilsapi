package model

type IsPalindromeRequest struct {
	A string `validate:"required" json:"a" example:"A man, a plan, a canal: Panama"`
}
