// go get github.com/gorilla/csrf

// curl -u admin:adminpassword -X POST http://localhost:8080/login -i
// curl -b "token=TOKEN http://localhost:8080/protected -i

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var users = map[string]string{
	"admin": "adminpassword",
	"user":  "userpassword",
}

func login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok || users[username] != password {
		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Role:     getRole(username),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Ошибка при генерации токена", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func getRole(username string) string {
	if username == "admin" {
		return "admin"
	}
	return "user"
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}

		tokenStr := c.Value

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Неверная подпись токена", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authorize(allowedRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie("token")
		tokenStr := c.Value

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid || claims.Role != allowedRole {
			http.Error(w, "Доступ запрещён", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func adminOnly(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Добро пожаловать, администратор!"))
}

func protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, это защищенный ресурс для аутентифицированных пользователей."))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", login).Methods("POST")

	r.Handle("/protected", authenticate(http.HandlerFunc(protected))).Methods("GET")
	r.Handle("/admin", authenticate(authorize("admin", http.HandlerFunc(adminOnly)))).Methods("GET")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
