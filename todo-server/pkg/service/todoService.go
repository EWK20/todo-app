package service

import (
	"log"
	"strconv"
	"todo-server/pkg/db"
	"todo-server/pkg/model"
	"todo-server/pkg/spec"
)

func GetAllTodos() ([]model.Todo, error) {

	var todos []model.Todo

	if err := db.DB.Find(&todos).Error; err != nil {
		log.Println("Could not get all todos")
		return nil, err
	}

	return todos, nil

}

func GetTodo(id int) (model.Todo, error) {

	var todo model.Todo

	if err := db.DB.Find(&todo, id).Error; err != nil {
		log.Println("Could not find todo")
		return todo, err
	}

	return todo, nil

}

func FilterTodos(item string, completed string) ([]model.Todo, error) {

	var todos []model.Todo

	if item == "" && completed == "" {

		if err := db.DB.Find(&todos).Error; err != nil {
			log.Println("Could not get todos")
			return nil, err
		}

	} else if item == "" && completed != "" {

		completedBool, err := strconv.ParseBool(completed)
		if err != nil {
			log.Println("Could not convert query to boolean")
		}
		completedSpec := spec.IsCompleted(completedBool)

		if err := db.DB.Scopes(completedSpec).Find(&todos).Error; err != nil {
			log.Println("Could not get completed todos")
			return nil, err
		}
	} else {
		itemSpec := spec.HasItemLike(item)

		completedBool, err := strconv.ParseBool(completed)
		if err != nil {
			log.Println("Could not convert query to boolean")
		}

		completedSpec := spec.IsCompleted(completedBool)

		if err := db.DB.Scopes(itemSpec, completedSpec).Find(&todos).Error; err != nil {
			log.Println("Could not get todos by spec")
			return nil, err
		}
	}

	return todos, nil
}

func CreateTodo(todo model.Todo) (model.Todo, error) {

	if err := db.DB.Create(&todo).Error; err != nil {
		log.Println("Could not create todo")
		return todo, err
	}
	return todo, nil
}

func ToggleStatus(id int) (model.Todo, error) {
	var todo model.Todo

	if err := db.DB.Find(&todo, id).Error; err != nil {
		log.Println("Could not toggle status")
		return todo, err
	}

	todo.Completed = !todo.Completed

	if err := db.DB.Model(&todo).Update("completed", todo.Completed).Error; err != nil {
		log.Println("Failed to toggle completed status")
	}

	return todo, nil
}

func EditTodo(id int, todoReq *model.Todo) (model.Todo, error) {

	var todoRes model.Todo

	if err := db.DB.Find(&todoRes, id).Error; err != nil {
		log.Println("Could not find todo")
		return todoRes, err
	}

	if err := db.DB.Model(&todoRes).Update("item", todoReq.Item).Error; err != nil {
		log.Println("Failed to update todo")
	}
	return todoRes, nil
}

func DeleteTodo(id int) (string, error) {
	var todo model.Todo

	if err := db.DB.Delete(todo, id).Error; err != nil {
		return "", err
	}
	return "Todo has been successfully deleted!", nil
}
