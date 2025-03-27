package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(username, password string) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user := User{
		Username: username,
		Password: string(passwordHash),
	}
	result := db.Where("username=?", username).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return false, nil
	}

	if err := db.Create(&user).Error; err != nil {
		return true, err
	}

	return true, nil
}

func CheckAuth(username, password string) (bool, error) {
	var user User
	result := db.Where("username=?", username).Find(&user)
	if result.RowsAffected < 1 {
		return false, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}
