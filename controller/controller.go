package controller

import (
	"net/http"

	"github.com/Digisata/digiutils"
	"github.com/Digisata/digiutilsapi/helper"
	"github.com/Digisata/digiutilsapi/model"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	Validate *validator.Validate
}

func NewController(validate *validator.Validate) *Controller {
	return &Controller{
		Validate: validate,
	}
}

func (c *Controller) AddBinary(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// swagger:route POST /add_binary addBinary
	//
	// produces:
	// - application/json
	// responses:
	//   200: SuccessResponse
	//   400: FailResponse
	requestBody := model.AddBinaryRequest{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := model.FailResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := model.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.AddBinary(requestBody.A, requestBody.B),
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *Controller) IsPalindrome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// swagger:route POST /is_palindrome isPalindrome
	//
	// produces:
	// - application/json
	// responses:
	//   200: SuccessResponse
	//   400: FailResponse
	requestBody := model.IsPalindromeRequest{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := model.FailResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := model.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.IsPalindrome(requestBody.A),
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *Controller) RomanToInt(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// swagger:route POST /roman_to_int romanToInt
	//
	// produces:
	// - application/json
	// responses:
	//   200: SuccessResponse
	//   400: FailResponse
	requestBody := model.RomanToIntRequest{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := model.FailResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := model.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.RomanToInt(requestBody.A),
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *Controller) ContainsDuplicate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// swagger:route POST /contains_duplicate containsDuplicate
	//
	// produces:
	// - application/json
	// responses:
	//   200: SuccessResponse
	//   400: FailResponse
	requestBody := model.ContainsDuplicateRequest{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := model.FailResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := model.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.ContainsDuplicate(requestBody.A),
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *Controller) IsPowerOfTwo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// swagger:route POST /is_power_of_two isPowerOfTwo
	//
	// produces:
	// - application/json
	// responses:
	//   200: SuccessResponse
	//   400: FailResponse
	requestBody := model.IsPowerOfTwoRequest{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := model.FailResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := model.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.IsPowerOfTwo(requestBody.A),
	}

	helper.WriteToResponseBody(w, webResponse)
}
