package controller

import (
	"net/http"

	"github.com/Digisata/digiutils"
	"github.com/Digisata/digiutilsapi/helper"
	"github.com/Digisata/digiutilsapi/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Validate *validator.Validate
}

func NewController(validate *validator.Validate) *Controller {
	return &Controller{
		Validate: validate,
	}
}

// AddBinary godoc
// @ID add-binary
// @Accept json
// @Produce json
// @Param AddBinaryRequest body model.AddBinaryRequest true "Add Binary"
// @Success 200 {object} model.SuccessResult
// @Failure 400 {object} model.BadRequestResult
// @Router /add_binary [post]
func (ctr *Controller) AddBinary(c echo.Context) error {
	requestBody := &model.AddBinaryRequest{}
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	err := ctr.Validate.Struct(requestBody)
	if err != nil {
		return helper.FailResponse(c, http.StatusBadRequest)
	}

	return helper.SuccessResponse(c, digiutils.AddBinary(requestBody.A, requestBody.B))
}

// IsPalindrome godoc
// @ID is-palindrome
// @Accept json
// @Produce json
// @Param IsPalindromeRequest body model.IsPalindromeRequest true "Is Palindrome"
// @Success 200 {object} model.SuccessResult
// @Failure 400 {object} model.BadRequestResult
// @Router /is_palindrome [post]
func (ctr *Controller) IsPalindrome(c echo.Context) error {
	requestBody := &model.IsPalindromeRequest{}
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	err := ctr.Validate.Struct(requestBody)
	if err != nil {
		return helper.FailResponse(c, http.StatusBadRequest)
	}

	return helper.SuccessResponse(c, digiutils.IsPalindrome(requestBody.A))
}

// RomanToInt godoc
// @ID roman-to-int
// @Accept json
// @Produce json
// @Param RomanToIntRequest body model.RomanToIntRequest true "Roman to Int"
// @Success 200 {object} model.SuccessResult
// @Failure 400 {object} model.BadRequestResult
// @Router /roman_to_int [post]
func (ctr *Controller) RomanToInt(c echo.Context) error {
	requestBody := &model.RomanToIntRequest{}
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	err := ctr.Validate.Struct(requestBody)
	if err != nil {
		return helper.FailResponse(c, http.StatusBadRequest)
	}

	return helper.SuccessResponse(c, digiutils.RomanToInt(requestBody.A))
}

// ContainsDuplicate godoc
// @ID contains-duplicate
// @Accept json
// @Produce json
// @Param ContainsDuplicateRequest body model.ContainsDuplicateRequest true "Contains Duplicate"
// @Success 200 {object} model.SuccessResult
// @Failure 400 {object} model.BadRequestResult
// @Router /contains_duplicate [post]
func (ctr *Controller) ContainsDuplicate(c echo.Context) error {
	requestBody := &model.ContainsDuplicateRequest{}
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	err := ctr.Validate.Struct(requestBody)
	if err != nil {
		return helper.FailResponse(c, http.StatusBadRequest)
	}

	return helper.SuccessResponse(c, digiutils.ContainsDuplicate(requestBody.A))
}

// IsPowerOfTwo godoc
// @ID is-power-of-two
// @Accept json
// @Produce json
// @Param IsPowerOfTwoRequest body model.IsPowerOfTwoRequest true "Is Power of Two"
// @Success 200 {object} model.SuccessResult
// @Failure 400 {object} model.BadRequestResult
// @Router /is_power_of_two [post]
func (ctr *Controller) IsPowerOfTwo(c echo.Context) error {
	requestBody := &model.IsPowerOfTwoRequest{}
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	err := ctr.Validate.Struct(requestBody)
	if err != nil {
		return helper.FailResponse(c, http.StatusBadRequest)
	}

	return helper.SuccessResponse(c, digiutils.IsPowerOfTwo(requestBody.A))
}
