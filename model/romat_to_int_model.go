package model

type RomanToIntRequest struct {
	// The roman numeral
	// 
	// required: true
    // example: III
	A string `validate:"required" json:"a"`
}

// swagger:parameters romanToInt
type RomanToIntBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RomanToIntRequest"
	//  required: true
	Body RomanToIntRequest `json:"body"`
 }
