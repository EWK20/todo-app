package main

import (
	"log"
	"net/http"
	"todo-server/pkg/controller"
	"todo-server/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	v1 := r.Group("/v1")
	api := v1.Group("/api")

	api.GET("/getTodoList", controller.GetAllTodos)
	api.GET("/getTodo/:id", controller.GetTodo)
	api.GET("/search", controller.FilterTodos)
	api.POST("/createTodo", controller.CreateTodo)
	api.PATCH("/toggleStatus/:id", controller.ToggleStatus)
	api.PATCH("/editTodo/:id", controller.EditTodo)
	api.DELETE("/deleteTodo/:id", controller.DeleteTodo)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
