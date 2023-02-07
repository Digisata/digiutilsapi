package app

import (
	"github.com/Digisata/digiutilsapi/controller"
	_ "github.com/Digisata/digiutilsapi/docs"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func NewRouter(c *controller.Controller) *echo.Echo {
	router := echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := router.Group("/api/v1")

	v1.POST("/add_binary", c.AddBinary)
	v1.POST("/is_palindrome", c.IsPalindrome)
	v1.POST("/roman_to_int", c.RomanToInt)
	v1.POST("/contains_duplicate", c.ContainsDuplicate)
	v1.POST("/is_power_of_two", c.IsPowerOfTwo)

	return router
}
