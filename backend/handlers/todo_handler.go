package handlers

import (
	"fmt"
	"net/http"

	"crud-app/db/sqlc"

	"github.com/gin-gonic/gin"
)

var _ TodoHandler = (*TodoHandlerImpl)(nil)

type TodoHandlerImpl struct {
	queries *sqlc.Queries
}

func NewTodoHandler(queries *sqlc.Queries) TodoHandler {
	return &TodoHandlerImpl{queries: queries}
}

func (h *TodoHandlerImpl) ListTodos(c *gin.Context) {
	todos, err := h.queries.ListTodos(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandlerImpl) CreateTodo(c *gin.Context) {
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

	todo, err := h.queries.CreateTodo(c, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandlerImpl) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if err := h.queries.DeleteTodo(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
