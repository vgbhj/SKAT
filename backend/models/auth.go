package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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
