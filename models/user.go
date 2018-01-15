package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int
	CreatedAt time.Time
	Username  string
	Email     string
	Password  string
	Bio       string
	Image     string
}

func EncryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}


func NewUser(email, username, password string) (*User, error) {
	if email == "" || username == "" || password == "" {
		return nil, fmt.Errorf("provided with empty fields")
	}
	return &User{
		Email:    email,
		Username: username,
		Password: EncryptPassword(password),
	}, nil
}

func (db *DB) CreateUser(user *User) error {
	u := User{}

	db.Find(&u, "email = ?", user.Email)
	if u != (User{}) {
		return fmt.Errorf("Email already exisits")
	}

	db.Find(&u, "username = ?", user.Username)
	if u != (User{}) {
		return fmt.Errorf("Username already exisits")
	}

	db.Create(user)

	return nil
}