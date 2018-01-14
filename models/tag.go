package models



type Tag struct {
	ID            uint
	Name          string `gorm:"unique"`
	TaggingsCount uint
	Articles      []Article `gorm:"many2many:taggings;"`
}
