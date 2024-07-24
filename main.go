package main

import (
	"fmt"
	"net/http"

	domain "github.com/dreamnajababy/todo-hugeman/domain"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todos API
// @version 1.0
// @description.markdown
// @termsOfService http://hugeman.com/

// @contact.name API Support
// @contact.url http://github.com/dreamnajababy
func main() {
	router := gin.Default()
	todoRepository := domain.InMemoryTodoRepository{}

	router.GET("/health", healthCheckHandler)
	router.GET("/todos", prepareGetTodoListHandler(domain.NewGetTodoListUseCase(&todoRepository)))
	router.GET("/todos/search", prepareSearchTodoHandler(domain.NewSearchTodoUseCase(&todoRepository)))
	router.PUT("/todos/:id", prepareUpdateTodoHandler(domain.NewUpdateTodoUsecase(&todoRepository)))
	router.POST("/todos", prepareCreateTodoHandler(domain.NewCreateTodoUseCase(&todoRepository)))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func healthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// prepareUpdateTodoHandler handles the updating of todos
// @Summary Update an existing todo
// @Description Update an existing todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param request body domain.UpdateTodoRequest true "Update Todo Request"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} gin.H
// @Router /todos [put]
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

// prepareSearchTodoHandler handles the searching of todos
// @Summary Search todos
// @Description Search todos based on filters
// @Tags todos
// @Accept json
// @Produce json
// @Param filter query string false "Filter"
// @Success 200 {array} domain.Todo
// @Failure 400 {object} gin.H
// @Router /search_todos [get]
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

// prepareCreateTodoHandler handles the creation of todos
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param request body domain.CreateTodoRequest true "Create Todo Request"
// @Success 201 {object} domain.Todo
// @Failure 400 {object} gin.H
// @Router /todos [post]
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

// prepareGetTodoListHandler handles the retrieval of the todo list
// @Summary Get list of todos
// @Description Get list of todos
// @Tags todos
// @Accept json
// @Produce json
// @Param filter query string false "Filter"
// @Success 200 {array} domain.Todo
// @Failure 400 {object} gin.H
// @Router /todos [get]
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
