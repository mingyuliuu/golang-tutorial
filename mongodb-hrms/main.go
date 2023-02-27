package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"mongodb-hrms/helpers"
)

func main() {
	if err := helpers.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", helpers.GetAllEmployees)
	app.Post("/employee", helpers.CreateEmployee)
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
}
