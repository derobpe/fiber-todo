package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // adding a logger middleware
	"github.com/derobpe/golang-fiber-basic-todo-app/routes" // import routes
)

func main() {
	
	// initiate fiber
	app := fiber.New()
	
	// adding a logger middleware
	app.Use(logger.New())

	// // give response when at
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"success":  true,
	// 		"message": "You are at the endpoint 😉",
	// 	})
	// })

	// setup routes
	setupRoutes(app)

	// listen to the server at 8000 port and catch the error if any
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}

	// separate function to handle all our routes
	func setupRoutes(app *fiber.App) {
		// moved from main method
		// give response when at /
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success":  true,
				"message": "You are at the endpoint 😉",
			})
		})
	
		// api group
		api := app.Group("/api")
	
		// give response when at /api
		api.Get("", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "You are at the api endpoint 😉",
			})
		})
	
		// connect todo routes
		routes.TodoRoute(api.Group("/todos"))
	}
