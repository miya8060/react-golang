package handlers

import "github.com/gin-gonic/gin"

type TodoLister interface {
	ListTodos(c *gin.Context)
}

type TodoCreator interface {
	CreateTodo(c *gin.Context)
}

type TodoDeleter interface {
	DeleteTodo(c *gin.Context)
}

type TodoUpdater interface {
	UpdateTodo(c *gin.Context)
}

// 全機能を必要とする場合のための集約インターフェース
type TodoHandler interface {
	TodoLister
	TodoCreator
	TodoDeleter
	TodoUpdater
}
