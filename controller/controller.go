package controller

import (
	"net/http"

	"github.com/Digisata/digiutils"
	"github.com/Digisata/digiutilsapi/helper"
	web "github.com/Digisata/digiutilsapi/model"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type controller struct {
	Validate *validator.Validate
}

func NewController(validate *validator.Validate) *controller {
	return &controller{
		Validate: validate,
	}
}

func (c *controller) AddBinary(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	requestBody := struct {
		A string `validate:"required" json:"a"`
		B string `validate:"required" json:"b"`
	}{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   nil,
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.AddBinary(requestBody.A, requestBody.B),
	}

	helper.WriteToResponseBody(w, webResponse)
}
