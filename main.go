package main

import (
	"fmt"
	"net/http"

	domain "github.com/dreamnajababy/todo-hugeman/domain"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	todoRepository := domain.InMemoryTodoRepository{}

	router.GET("/health", healthCheckHandler)
	router.GET("/todos", prepareGetTodoListHandler(domain.NewGetTodoListUseCase(&todoRepository)))
	router.GET("/todos/search", prepareSearchTodoHandler(domain.NewSearchTodoUseCase(&todoRepository)))
	router.PUT("/todos/:id", prepareUpdateTodoHandler(domain.NewUpdateTodoUsecase(&todoRepository)))
	router.POST("/todos", prepareCreateTodoHandler(domain.NewCreateTodoUseCase(&todoRepository)))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("8080")
}

func healthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func prepareUpdateTodoHandler(updateTodoUseCase domain.UpdateTodoUseCase) func(c *gin.Context) {
	updateTodoHandler := func(c *gin.Context) {
		var request domain.UpdateTodoRequest

		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := updateTodoUseCase.Execute(request)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, todo)
	}
	return updateTodoHandler
}

func prepareSearchTodoHandler(searchTodoUseCase domain.SearchTodoUseCase) func(c *gin.Context) {
	searchTodoHandler := func(c *gin.Context) {
		var request domain.SearchTodoRequest

		if err := c.ShouldBindQuery(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todoList := searchTodoUseCase.Execute(request)
		c.IndentedJSON(http.StatusOK, todoList)
	}
	return searchTodoHandler
}

func prepareCreateTodoHandler(createTodoUseCase domain.CreateTodoUseCase) func(c *gin.Context) {
	createTodoHandler := func(c *gin.Context) {
		var request domain.CreateTodoRequest

		// Bind incoming JSON to newTodo
		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert CreateTodoRequest to Todo
		todo, err := createTodoUseCase.Execute(request)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, todo)
	}
	return createTodoHandler
}

func prepareGetTodoListHandler(getTodoListUseCase domain.GetTodoListUseCase) func(c *gin.Context) {
	getTodoListHandler := func(c *gin.Context) {
		var request domain.GetTodoListRequest

		if err := c.ShouldBindQuery(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(request)
		todoList := getTodoListUseCase.Execute(request)

		fmt.Print(todoList)
		c.IndentedJSON(http.StatusOK, todoList)
	}
	return getTodoListHandler
}
