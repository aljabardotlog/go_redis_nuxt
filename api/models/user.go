package models

type User struct {
	Base
	ID       uint `gorm:"column:id;primary_key;auto_increment"`
	Username string
	Password string
}
