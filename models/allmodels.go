package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Address string
	UserName    string
	Password    string
}


type Book struct {
	gorm.Model
	Title   string
	Publisher string
	Author string
	Year string
}