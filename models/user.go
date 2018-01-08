package models

import "time"

type User struct {
	ID        int
	CreatedAt time.Time
	Username  string
	Email     string
	Password  string
	Bio       string
	Image     string
}

