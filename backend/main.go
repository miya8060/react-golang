package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"crud-app/db"
	"crud-app/db/sqlc"
)

func main() {
	database := db.NewDB()
	defer database.Close()

	queries := sqlc.New(database)

	r := gin.Default()

	// CORSミドルウェアの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// テスト用エンドポイント
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Hello from Go backend!",
		})
	})

	// TODOのエンドポイント
	r.GET("/api/todos", func(c *gin.Context) {
		todos, err := queries.ListTodos(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/api/todos", func(c *gin.Context) {
		// Content-Typeの検証
		if c.GetHeader("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
			return
		}

		var req struct {
			Title string `json:"title" binding:"required,min=1"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := queries.CreateTodo(c, req.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, todo)
	})

	r.DELETE("/api/todos/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		var id int32
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
			return
		}
		if err := queries.DeleteTodo(c, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})

	// サーバーの起動
	r.Run(":8080")
}
