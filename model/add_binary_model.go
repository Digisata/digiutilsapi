package model

type AddBinaryRequest struct {
	// The binary string
	// 
	// required: true
    // example: 1101
	A string `validate:"required" json:"a"`

	// The binary string
	// 
	// required: true
    // example: 11
	B string `validate:"required" json:"b"`
}

// swagger:parameters addBinary
type AddBinaryBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/AddBinaryRequest"
	//  required: true
	Body AddBinaryRequest `json:"body"`
 }
