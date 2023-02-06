package app

import (
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(c *controller.Controller) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/add_binary", c.AddBinary)
	router.POST("/api/v1/is_palindrome", c.IsPalindrome)
	router.POST("/api/v1/roman_to_int", c.RomanToInt)
	router.POST("/api/v1/contains_duplicate", c.ContainsDuplicate)
	router.POST("/api/v1/is_power_of_two", c.IsPowerOfTwo)

	return router
}
