package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int    `gorm:"primary"` //row_id
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	CidNumber   bool   `json:"cid_number"`
}

func (User) TableName() string {
	return "user"
}

func HashPassword(password string) string {
	// cost := 7 // Min: 4, Max: 31
	// bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	bytes := md5.Sum([]byte(password))

	// return string(bytes)
	return fmt.Sprintf("%x", bytes)
}

func CheckPasswordHash(hashedPassword, password string) error {
	// return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	providedPassword := HashPassword(password)
	if providedPassword == hashedPassword {
		return nil
	} else {
		return errors.New("error")
	}
}
