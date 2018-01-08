package models

import "time"

type Article struct {
	ID             int
	Slug           string
	Title          string
	Description    string
	Body           string
	User           User
	UserID         int
	//Tags           []Tag `gorm:"many2many:taggings;"`
	//Favorites      []Favorite
	FavoritesCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
