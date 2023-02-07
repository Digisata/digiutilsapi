package model

type IsPowerOfTwoRequest struct {
	// The int to be checked for
	// 
	// required: true
    // example: 12
	A int `validate:"required" json:"a"`
}

// swagger:parameters isPowerOfTwo
type IsPowerOfTwoBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/IsPowerOfTwoRequest"
	//  required: true
	Body IsPowerOfTwoRequest `json:"body"`
 }