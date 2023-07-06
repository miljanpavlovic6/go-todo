package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	instana "github.com/instana/go-sensor"
)

func main() {

	c := instana.InitCollector(&instana.Options{
		Service: "my-go-app",
	})

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ok")
	}

	http.HandleFunc("/api", instana.TracingHandlerFunc(c.LegacySensor(), "/api", handler))
	//
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Test handler
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Test App is running")
	})

	log.Fatal(app.Listen(":5000"))
}
