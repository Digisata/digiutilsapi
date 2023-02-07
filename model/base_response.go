package model

// swagger:model SuccessResponse
type SuccessResponse struct {
	// the http status code
	//
	// required: true
	// example: 200
	Code int `json:"code"`

	// the http status
	//
	// required: true
	// example: Ok
	Status string `json:"status"`

	// the actual data
	//
	// required: true
	// example: any
	Data any `json:"data"`
}

// swagger:model FailResponse
type FailResponse struct {
	// the http status code
	//
	// required: true
	// example: 400
	Code int `json:"code"`

	// the http status
	//
	// required: true
	// example: BAD REQUEST
	Status string `json:"status"`
}
