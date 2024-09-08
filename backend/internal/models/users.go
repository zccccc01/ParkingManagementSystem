package models

type User struct {
	UserID   int `gorm:"primaryKey"`
	Username string
	Password string
	Tel      string
}
