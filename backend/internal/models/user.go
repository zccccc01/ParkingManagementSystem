package models

type User struct {
	UserID   int    `gorm:"column:UserID;primaryKey"`
	UserName string `gorm:"column:Username;size:255;not null"`
	Password string `gorm:"column:Password;size:255;not null"`
	Tel      string `gorm:"column:Tel;size:36"`
}

func (u *User) TableName() string {
	return "users"
}
