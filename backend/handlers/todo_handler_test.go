package handlers

import (
	"bytes"
	"context"
	"crud-app/db/sqlc"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueries は sqlc.Queries のモック
type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) ListTodos(ctx context.Context) ([]sqlc.Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]sqlc.Todo), args.Error(1)
}

func (m *MockQueries) CreateTodo(ctx context.Context, title string) (sqlc.Todo, error) {
	args := m.Called(ctx, title)
	return args.Get(0).(sqlc.Todo), args.Error(1)
}

func (m *MockQueries) DeleteTodo(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockQueries) GetTodo(ctx context.Context, id int32) (sqlc.Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(sqlc.Todo), args.Error(1)
}

func TestListTodos(t *testing.T) {
	mockQueries := new(MockQueries)
	handler := NewTodoHandler(mockQueries)

	// テストケースの設定
	todos := []sqlc.Todo{{ID: 1, Title: "Test Todo"}}
	mockQueries.On("ListTodos", mock.Anything).Return(todos, nil)

	// HTTPリクエストの設定
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// ハンドラーの実行
	handler.ListTodos(c)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	var response []sqlc.Todo
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, todos, response)
}

func TestCreateTodo(t *testing.T) {
	mockQueries := new(MockQueries)
	handler := NewTodoHandler(mockQueries)

	// テストケースの設定
	newTodo := sqlc.Todo{ID: 1, Title: "New Todo"}
	mockQueries.On("CreateTodo", mock.Anything, "New Todo").Return(newTodo, nil)

	// HTTPリクエストの設定
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewBufferString(`{"title":"New Todo"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// ハンドラーの実行
	handler.CreateTodo(c)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	var response sqlc.Todo
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newTodo, response)
}

func TestDeleteTodo(t *testing.T) {
	mockQueries := new(MockQueries)
	handler := NewTodoHandler(mockQueries)

	// テストケースの設定
	mockQueries.On("DeleteTodo", mock.Anything, int32(1)).Return(nil)

	// HTTPリクエストの設定
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	// ハンドラーの実行
	handler.DeleteTodo(c)

	// アサーション
	assert.Equal(t, http.StatusNoContent, w.Code)
}
