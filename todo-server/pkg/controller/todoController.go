package controller

import (
	"log"
	"net/http"
	"strconv"
	"todo-server/pkg/model"
	"todo-server/pkg/service"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(ctx *gin.Context) {

	todos, err := service.GetAllTodos()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Could not find todo"})
	}

	ctx.IndentedJSON(http.StatusFound, todos)
}

func GetTodo(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalln("Could not convert string to int")
	}

	todo, err := service.GetTodo(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Could not find todo"})
	}
	ctx.IndentedJSON(http.StatusFound, todo)
}

func FilterTodos(ctx *gin.Context) {

	item := ctx.Query("item")
	completed := ctx.Query("completed")

	todos, err := service.FilterTodos(item, completed)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Could not find the todos"})
	}
	ctx.IndentedJSON(http.StatusOK, todos)
}

func CreateTodo(ctx *gin.Context) {

	var todo model.Todo

	err := ctx.BindJSON(&todo)
	if err != nil {
		log.Fatalln("Could not retrieve JSON data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	newTodo, err := service.CreateTodo(todo)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Could not create todo"})
	}
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func ToggleStatus(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalln("Could not convert string to int")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	todo, err := service.ToggleStatus(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"err": "Could not toggle status"})
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func EditTodo(ctx *gin.Context) {

	var todoReq *model.Todo

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalln("Could not convert string to int")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	err = ctx.BindJSON(&todoReq)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	todo, err := service.EditTodo(id, todoReq)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Failed to edit todo"})
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalln("Could not convert string to int")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	msg, err := service.DeleteTodo(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"err": "Todo could not be deleted!"})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"msg": msg})
}
