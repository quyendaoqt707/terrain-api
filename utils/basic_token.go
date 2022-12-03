package utils

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

func GenBase64Token(email string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(email)))
}

func ExtractBase64Token(token string) (string, error) {
	email, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	return string(email), nil
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
