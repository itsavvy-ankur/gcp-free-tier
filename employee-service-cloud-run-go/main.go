package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//Employee Service Domain
type Employee struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

func main() {
	app := fiber.New()
	// Default middleware config
	app.Use(logger.New())
	app.Get("/employee", func(c *fiber.Ctx) error {
		return c.SendString(employeeHandler(c))
	})

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	// Start server on port 8080
	app.Listen(":" + port)
}

//Employee Handler
func employeeHandler(c *fiber.Ctx) string {

	//Employee
	e := Employee{
		Name:    "John",
		Email:   "john.doe@test.com",
		Country: "USA"}

	//Marshall Employee
	data, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
