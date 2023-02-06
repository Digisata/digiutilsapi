package main

import (
	"net/http"

	"github.com/Digisata/digiutilsapi/app"
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	controller := controller.NewController(validate)
	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
