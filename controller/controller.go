package controller

import (
	"net/http"

	"github.com/Digisata/digiutils"
	"github.com/Digisata/digiutilsapi/helper"
	web "github.com/Digisata/digiutilsapi/model"
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
	requestBody := struct {
		A string `validate:"required" json:"a"`
		B string `validate:"required" json:"b"`
	}{}
	helper.ReadFromRequestBody(r, &requestBody)
	err := c.Validate.Struct(requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   nil,
		}

		helper.WriteToResponseBody(w, webResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   digiutils.AddBinary(requestBody.A, requestBody.B),
	}

	helper.WriteToResponseBody(w, webResponse)
}
