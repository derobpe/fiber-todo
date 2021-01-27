package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/derobpe/golang-fiber-basic-todo-app/controllers"
)

// first route to get all todos
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)

	// other routes to new controllers
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}