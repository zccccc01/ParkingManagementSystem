package models

type User struct {
	UserID   int    `gorm:"column:UserID;primaryKey"`
	Username string `gorm:"column:Username;size:255;not null"`
	Password string `gorm:"column:Password;size:255;not null"`
	Tel      string `gorm:"column:Tel;size:36"`
}
