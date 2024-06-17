package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	IsAdmin  bool
	Username string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
	Nama     string `gorm:"type:varchar(50)"`
	// BirthDate time.Time `gorm:"type:date"`
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) *UserModel {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(username string, password string) (User, error) {
	var result User
	err := um.db.Where("username = ? AND password = ?", username, password).First(&result).Error
	if err != nil {
		return User{}, err
	}
	return result, nil
}
