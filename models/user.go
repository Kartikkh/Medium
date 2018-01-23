package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

)

type User struct {
	ID         int
	CreatedAt  time.Time
	Username   string
	Email      string
	Password   string
	Bio        string
	Image      string
	Following  []User

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
	if u.Email != "" {
		return fmt.Errorf("email already exits")
	}

	db.Find(&u, "username = ?", user.Username)
	if u.Username != ""  {
		return fmt.Errorf("username already exits")
	}

	db.Create(user)

	return nil
}


func (db *DB) FindUserByEmail(email string) (*User, error) {
	u := User{}
	db.Find(&u, "email = ?", email)
	if u.Email == "" {
		return nil, fmt.Errorf("No user found with email: ", email)
	}
	return &u, nil
}

func (u *User) MatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (db *DB) FindUserByUserName(username string) (*User, error) {
	u := User{}
	db.Find(&u, "username = ?", username)
	if u.Username ==  "" {
		return nil, fmt.Errorf("No user found with userame: ", username)
	}
	return &u, nil
}


func (db *DB) IsFollowing( userCheck *User, LoginUser *User) bool  {
	userPr := &User{}
	for  _ , userPr =  range LoginUser.Following {
		if userPr.Username == userCheck.Username {
			return true
		}
	}

	return false
}

