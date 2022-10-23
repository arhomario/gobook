package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) SetToken(token string) {
	u.Token = token
}
