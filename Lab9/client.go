// curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"username\":\"admin\", \"password\":\"password\"}"
// curl -X GET http://localhost:8080/users -H "Authorization: Bearer TOKEN"

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseURL     = "http://localhost:8080/users"
	loginURL    = "http://localhost:8080/login" // Эндпоинт для авторизации
	contentType = "application/json"
	authHeader  = "Authorization"
)

var sessionToken string

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var token string
var client = &http.Client{}

func main() {
	data := []byte(`{"example": "data"}`)
	buffer := bytes.NewBuffer(data)
	response, err := makeAuthorizedRequest("POST", "http://example.com/api", buffer)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	if !login() {
		fmt.Println("Не удалось войти. Программа завершена.")
		return
	}

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Показать всех пользователей")
		fmt.Println("2. Добавить нового пользователя")
		fmt.Println("3. Обновить данные пользователя")
		fmt.Println("4. Удалить пользователя")
		fmt.Println("5. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showUsers()
		case 2:
			addUser()
		case 3:
			updateUser()
		case 4:
			deleteUser()
		case 5:
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func login() bool {
	fmt.Println("Введите логин:")
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	loginData := map[string]string{"username": username, "password": password}
	data, _ := json.Marshal(loginData)

	resp, err := http.Post(loginURL, contentType, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Ошибка авторизации: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка авторизации. Проверьте логин и пароль.")
		return false
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Ошибка при обработке ответа авторизации: %v\n", err)
		return false
	}

	sessionToken = result["token"]
	token = sessionToken // Установите глобальный токен для использования в запросах

	fmt.Println("Авторизация успешна! Токен:", sessionToken) // выводим токен для проверки
	return true
}

// showUsers - выводит список пользователей
func showUsers() {
	resp, err := makeAuthorizedRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	var result struct {
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		Results []User `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Ошибка при обработке данных: %v\n", err)
		return
	}

	fmt.Println("Список пользователей:")
	for _, user := range result.Results {
		fmt.Printf("ID: %d, Имя: %s, Возраст: %d\n", user.ID, user.Name, user.Age)
	}
}

// addUser - добавляет нового пользователя
func addUser() {
	var user User
	fmt.Print("Введите имя: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&user.Age)

	data, _ := json.Marshal(user)
	resp, err := makeAuthorizedRequest(http.MethodPost, baseURL, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Ошибка создания пользователя:", err)
		return
	}
	defer resp.Body.Close()

	printResponse(resp)
}

// updateUser - обновляет данные пользователя
func updateUser() {
	var id int
	fmt.Print("Введите ID пользователя для обновления: ")
	fmt.Scanln(&id)

	var user User
	fmt.Print("Введите новое имя: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Введите новый возраст: ")
	fmt.Scanln(&user.Age)

	data, _ := json.Marshal(user)
	resp, err := makeAuthorizedRequest(http.MethodPut, fmt.Sprintf("%s/%d", baseURL, id), bytes.NewBuffer(data))
	if err != nil {
		log.Println("Ошибка обновления пользователя:", err)
		return
	}
	defer resp.Body.Close()

	printResponse(resp)
}

// deleteUser - удаляет пользователя
func deleteUser() {
	var id int
	fmt.Print("Введите ID пользователя для удаления: ")
	fmt.Scanln(&id)

	resp, err := makeAuthorizedRequest(http.MethodDelete, fmt.Sprintf("%s/%d", baseURL, id), nil)
	if err != nil {
		log.Println("Ошибка удаления пользователя:", err)
		return
	}
	defer resp.Body.Close()

	printResponse(resp)
}

func makeAuthorizedRequest(method, url string, payload *bytes.Buffer) (*http.Response, error) {
	var req *http.Request
	var err error

	if payload != nil {
		req, err = http.NewRequest(method, url, payload)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Установка токена в заголовок Authorization
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	return client.Do(req)
}

// printResponse - обрабатывает и выводит ответ от сервера
func printResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка чтения ответа: %v\n", err)
		return
	}
	fmt.Println("Ответ сервера:", string(body))
}
