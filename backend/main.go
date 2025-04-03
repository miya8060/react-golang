package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"crud-app/db"
	"crud-app/db/sqlc"
	"crud-app/handlers"
)

func setupRouter(queries *sqlc.Queries) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	var todoHandler handlers.TodoHandler = handlers.NewTodoHandler(queries)

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Hello from Go backend!",
			})
		})

		todos := api.Group("/todos")
		{
			todos.GET("", todoHandler.ListTodos)
			todos.POST("", todoHandler.CreateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	return r
}

func main() {
	database := db.NewDB()
	defer database.Close()

	queries := sqlc.New(database)
	r := setupRouter(queries)
	r.Run(":8080")
}
