package entity

import (
	"github.com/alcir-junior-caju/study-go-expert-client-apis/pkg/valueObject"
	"golang.org/x/crypto/bcrypt"
)

type UserStruct struct {
	ID       valueObject.IDType `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
}

func User(name, email, password string) (*UserStruct, error) {
	passwordHash, passwordHashError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if passwordHashError != nil {
		panic(passwordHashError)
	}
	return &UserStruct{
		ID:       valueObject.ID(),
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}, nil
}

func (user *UserStruct) ValidatePassword(password string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return error == nil
}
