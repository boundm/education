package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	jwtKey = []byte("my_secret_key")
	store  = sessions.NewCookieStore([]byte("session_key"))
)

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
		Role:     "user", // задайте роль в зависимости от пользователя
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

	w.Write([]byte("Login successful"))
}

func protected(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Привет, это защищенный ресурс для аутентифицированных пользователей."))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/protected", protected).Methods("GET")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
