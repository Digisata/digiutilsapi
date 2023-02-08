package main

import (
	"fmt"
	"github.com/Digisata/digiutilsapi/spec"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/Digisata/digiutilsapi/app"
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
)

// @title           Digiutilsapi Specification
// @description     The purpose of this application is to provide an API for handy tools to solve trivial daily life problems.
// @termsOfService  http://swagger.io/terms/
func main() {
	host := os.Getenv("DOCS_HOST")
	port := os.Getenv("PORT")
	schemes := []string{"https"}

	env := os.Getenv("ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		port = os.Getenv("PORT")
		host = fmt.Sprintf("%s:%s", os.Getenv("DOCS_HOST"), port)
		schemes = []string{"http"}
	}

	spec.SwaggerInfo.Version = os.Getenv("DOCS_VERSION")
	spec.SwaggerInfo.Host = host
	spec.SwaggerInfo.BasePath = os.Getenv("DOCS_BASE_PATH")
	spec.SwaggerInfo.Schemes = schemes

	validate := validator.New()
	controller := controller.NewController(validate)
	router := app.NewRouter(controller)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", port)))
}
