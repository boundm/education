package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var jwtKey = []byte("my_secret_key")
var csrfKey = []byte("my_csrf_key")
var store = sessions.NewCookieStore([]byte("session_key"))

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var users = map[string]string{
	"admin": "adminpassword",
	"user":  "userpassword",
}

func getRole(username string) string {
	if username == "admin" {
		return "admin"
	}
	return "user"
}

func login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok || users[username] != password {
		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session-name")
	session.Values["username"] = username
	session.Values["role"] = getRole(username)
	session.Save(r, w)

	csrfToken := csrf.Token(r) // Получаем CSRF-токен

	http.SetCookie(w, &http.Cookie{
		Name:     "_gorilla_csrf",
		Value:    csrfToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Установите на true в производственной среде
	})

	w.Write([]byte("CSRF token: " + csrfToken)) // Добавляем вывод CSRF-токена
}

func authenticateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil || session.Values["username"] == nil {
			http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authorize(allowedRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil || session.Values["role"] != allowedRole {
			http.Error(w, "Доступ запрещён", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, это защищенный ресурс для аутентифицированных пользователей."))
}

func adminOnly(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Добро пожаловать, администратор!"))
}

func main() {
	r := mux.NewRouter()

	CSRF := csrf.Protect(csrfKey, csrf.Secure(false))

	r.HandleFunc("/login", login).Methods("POST")
	r.Handle("/protected", authenticateSession(http.HandlerFunc(protected))).Methods("GET")
	r.Handle("/admin", authenticateSession(authorize("admin", http.HandlerFunc(adminOnly)))).Methods("GET")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", CSRF(r)))
}
