package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	return r
}

func TestMain(m *testing.M) {
	initDB()             // Инициализация базы данных
	defer dbPool.Close() // Закрытие при завершении

	os.Exit(m.Run()) // Запуск тестов
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	user := User{Name: "test", Age: 30}
	userJSON, _ := json.Marshal(user)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", res.Code)
	}

	var createdUser User
	json.Unmarshal(res.Body.Bytes(), &createdUser)

	if createdUser.Name != "TestUser" {
		t.Errorf("Expected user name to be 'TestUser', got %s", createdUser.Name)
	}
	if createdUser.Age != 30 {
		t.Errorf("Expected user age to be 30, got %d", createdUser.Age)
	}
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?page=1&limit=10", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.Code)
	}

	var usersResponse struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		Results []User `json:"results"`
	}
	json.Unmarshal(res.Body.Bytes(), &usersResponse)

	if usersResponse.Page != 1 {
		t.Errorf("Expected page to be 1, got %d", usersResponse.Page)
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupRouter()

	user := User{Name: "UpdatedUser", Age: 35}
	userJSON, _ := json.Marshal(user)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	router := setupRouter()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.Code)
	}
}
