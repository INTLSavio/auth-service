package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Password string `json:"password"`
}

func NewUser(name, email, phone, address, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Phone: phone,
		Password: string(hash),
	}, nil
}
