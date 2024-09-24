package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"gorm.io/gorm"

	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.StandardClaims
}

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println("Ошибка парсинга файлов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title           string
		IsAuthenticated bool
	}{
		Title:           "Вход",
		IsAuthenticated: false,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка execute: ", err) //TODO
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			renderLoginPageWithError(w, "Неверные учетные данные")
			return
		}
		log.Println("Ошибка поиска пользователя:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Admin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderLoginPageWithError(w http.ResponseWriter, errorMessage string) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблонов:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title        string
		ErrorMessage string
	}{
		Title:        "Вход",
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")

	if err != nil {
		log.Println("Ошибка парсинга файлов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title        string
		ErrorMessage string
	}{
		Title:        "Регистрация",
		ErrorMessage: "",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка execute: ", err) //TODO
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return // Добавьте return здесь
	}
}

func renderRegisterPageWithError(w http.ResponseWriter, errorMessage string) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблонов:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title        string
		ErrorMessage string
	}{
		Title:        "Регистрация",
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	var existingUser models.User

	result := database.DB.Where("email = ?", email).First(&existingUser)
	if result.Error == nil {
		renderRegisterPageWithError(w, "Пользователь с таким email уже существует")
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		log.Println("Ошибка поиска пользователя:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		IsAdmin:  false,
	}

	result = database.DB.Create(&user)
	if result.Error != nil {
		log.Println("Ошибка создания пользователя:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Удаляем токен
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	// Перенаправляем на главную страницу после выхода
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
