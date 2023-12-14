package models

type User struct {
	Id int `gorm:"primaryKey" json=id`
	Fullname string `gorm:"varchar(200)" json=fullname`
	Username string `gorm:"varchar(200)" json=username`
	Password string `gorm:"varchar(200)" json=password`
}