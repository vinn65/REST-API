package main

import (
	
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct{
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = [] todo{
	{ID : "1", Item: "Clean Room", Completed: false},
	{ID : "2", Item: "Finish assignment 2", Completed: false},
	{ID : "1", Item: "Watch two episodes.", Completed: false},

}
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
	var newTodo todo
	
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}



func getTodoById(id string) (* todo, error){
	for i, t := range todos {
		if t.ID == id{
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found!!")
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message":"Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func patchTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message":"Todo not found"})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main(){
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/todos", getTodos)
	router.GET("/todos:id", getTodo)
	router.PATCH("/todos:id", patchTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")

}