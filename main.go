package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Digisata/digiutilsapi/app"
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	port := os.Getenv("PORT")

	validate := validator.New()
	controller := controller.NewController(validate)
	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	fmt.Printf("running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
