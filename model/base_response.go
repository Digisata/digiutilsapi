package model

type SuccessResult struct {
	Code   int    `json:"code" example:"200"`
	Status string `json:"status" example:"Ok"`
	Data   any    `json:"data"`
}

type BadRequestResult struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"BAD REQUEST"`
}
