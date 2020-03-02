package models

import ("github.com/jinzhu/gorm"
"time")

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Address string
	UserName    string
	Password    string
}

type Category struct {
	gorm.Model
	Name   string
	Description string
	Book 		[]Book `gorm:ForeignKey"CategoryID"`
}
type Book struct {
	gorm.Model
	Title   string
	Publisher string
	Author string
	Year string
	CategoryID uint
	PublicationID uint
}

type IssuedBook struct {
	gorm.Model
	UserID   uint
	BookID uint
	IssueDate *time.Time
}
type ReturnedBook struct {
	gorm.Model
	UserID   uint
	BookID uint
	ReturnDate *time.Time
}
type Publication struct {
	gorm.Model
	Name   string
	Address string
	Phone string
	Email string
	Book []Book
}