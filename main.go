package main

import (
	"net/http"

	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	validate := validator.New()

	controller := controller.NewController(validate)

	router.POST("/api/v1/add_binary", controller.AddBinary)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
