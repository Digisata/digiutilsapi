package main

import (
	"fmt"
	"os"

	"github.com/Digisata/digiutilsapi/app"
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// @title           Digiutilsapi Specification
// @version         1.0.0
// @description     The purpose of this application is to provide an API for handy tools to solve trivial daily life problems.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:3000
// @BasePath  /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	port := os.Getenv("PORT")

	validate := validator.New()
	controller := controller.NewController(validate)
	router := app.NewRouter(controller)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", port)))
}
