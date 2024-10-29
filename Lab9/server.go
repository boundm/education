package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=100"`
	Age  int    `json:"age" validate:"required,gt=0,lt=120"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var dbPool *pgxpool.Pool
var validate *validator.Validate
var jwtSecret = []byte("my_secret_key")

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

func initDB() {
	var err error
	dbURL := "postgres://postgres:516725@localhost:5432/userdb"
	dbPool, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to the database!")
}

func main() {
	initDB()
	defer dbPool.Close()

	validate = validator.New()

	r := gin.Default()

	// Маршрут для логина
	r.POST("/login", loginHandler)

	// Маршруты для пользователей, защищенные JWT
	protected := r.Group("/", authMiddleware)
	{
		protected.GET("/users", getUsers)
		protected.POST("/users", createUser)
		protected.PUT("/users/:id", updateUser)
		protected.DELETE("/users/:id", deleteUser)
	}

	r.Run() // http://localhost:8080
}

// Handler для логина
func loginHandler(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid login data")
		return
	}

	if creds.Username != "admin" || creds.Password != "password" {
		handleError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Could not generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Middleware для аутентификации
func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		handleError(c, http.StatusUnauthorized, "Authorization token not provided")
		c.Abort()
		return
	}

	// Убираем "Bearer " из строки токена
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		handleError(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	c.Next()
}

func getUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	nameFilter := c.Query("name")
	ageFilter := c.Query("age")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		handleError(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		handleError(c, http.StatusBadRequest, "Invalid limit number")
		return
	}

	offset := (page - 1) * limit

	query := "SELECT id, name, age FROM users WHERE 1=1"
	var args []interface{}
	var conditions []string

	if nameFilter != "" {
		conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+nameFilter+"%")
	}

	if ageFilter != "" {
		age, err := strconv.Atoi(ageFilter)
		if err != nil {
			handleError(c, http.StatusBadRequest, "Invalid age parameter")
			return
		}
		conditions = append(conditions, "age = $"+strconv.Itoa(len(args)+1))
		args = append(args, age)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT %d OFFSET %d", limit, offset)

	rows, err := dbPool.Query(context.Background(), query, args...)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			handleError(c, http.StatusInternalServerError, "Failed to scan user")
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"results": users,
	})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if validate == nil {
		handleError(c, http.StatusInternalServerError, "Validator is not initialized")
		return
	}

	log.Printf("Creating user: %+v\n", newUser)

	if err := validate.Struct(&newUser); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	var id int
	err := dbPool.QueryRow(context.Background(), "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Age).Scan(&id)
	if err != nil {
		log.Printf("Failed to insert user: %v\n", err) // Логирование ошибки
		handleError(c, http.StatusInternalServerError, "Failed to insert user")
		return
	}

	newUser.ID = id
	c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if validate == nil {
		handleError(c, http.StatusInternalServerError, "Validator is not initialized")
		return
	}

	log.Printf("Updating user ID %d: %+v\n", id, updatedUser)

	if err := validate.Struct(&updatedUser); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = dbPool.Exec(context.Background(), "UPDATE users SET name=$1, age=$2 WHERE id=$3", updatedUser.Name, updatedUser.Age, id)
	if err != nil {
		log.Printf("Failed to update user: %v\n", err)
		handleError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	updatedUser.ID = id
	c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	_, err = dbPool.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
