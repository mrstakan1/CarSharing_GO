package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	password := "testpassword"

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Не удалось хешировать пароль: %v", err)
	}

	// Проверяем хеш с правильным паролем
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		t.Errorf("Пароль должен совпадать с хешем")
	}

	// Проверяем хеш с неправильным паролем
	wrongPassword := "wrongpassword"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(wrongPassword))
	if err == nil {
		t.Errorf("Хеш не должен совпадать с неправильным паролем")
	}
}
