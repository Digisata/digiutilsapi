package model

type ContainsDuplicateRequest struct {
	// The string to be checked for
	//
	// required: true
	// example: [1, 2, 3, 1]
	A []int `validate:"required" json:"a"`
}

// swagger:parameters containsDuplicate
type ContainsDuplicateBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ContainsDuplicateRequest"
	//  required: true
	Body ContainsDuplicateRequest `json:"body"`
}
