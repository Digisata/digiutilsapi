package app

import (
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(c *controller.Controller) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/add_binary", c.AddBinary)

	return router
}
