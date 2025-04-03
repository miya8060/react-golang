package main

import (
	"crud-app/db/sqlc"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// テスト用のルーターを設定
	queries := &sqlc.Queries{}
	router := setupRouter(queries)

	// テストリクエストを作成
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)

	// リクエストを実行
	router.ServeHTTP(w, req)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Hello from Go backend!")
}

func TestCORSMiddleware(t *testing.T) {
	// テスト用のルーターを設定
	queries := &sqlc.Queries{}
	router := setupRouter(queries)

	// テストリクエストを作成
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/api/todos", nil)
	req.Header.Set("Origin", "http://localhost:5173")

	// リクエストを実行
	router.ServeHTTP(w, req)

	// アサーション
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "http://localhost:5173", w.Header().Get("Access-Control-Allow-Origin"))
}
