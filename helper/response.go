package helper

import (
	"net/http"

	"github.com/Digisata/digiutilsapi/model"
	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, model.SuccessResult{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   data,
	})
}

func FailResponse(c echo.Context, code int) error {
	switch code {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, model.BadRequestResult{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		})
	}
	return nil
}
