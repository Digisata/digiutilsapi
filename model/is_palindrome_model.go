package model

type IsPalindromeRequest struct {
	// The string to be checked for
	//
	// required: true
	// example: A man, a plan, a canal: Panama
	A string `validate:"required" json:"a"`
}

// swagger:parameters isPalindrome
type IsPalindromeBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/IsPalindromeRequest"
	//  required: true
	Body IsPalindromeRequest `json:"body"`
}
