package database

import (
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"gorm.io/gorm"
)

type UserDatabaseStruct struct {
	Connection *gorm.DB
}

func UserDatabase(connection *gorm.DB) *UserDatabaseStruct {
	return &UserDatabaseStruct{Connection: connection}
}

func (userDatabase *UserDatabaseStruct) Create(user *entity.UserStruct) error {
	return userDatabase.Connection.Create(user).Error
}

func (userDatabase *UserDatabaseStruct) FindByEmail(email string) (*entity.UserStruct, error) {
	var user entity.UserStruct
	if errorUser := userDatabase.Connection.Where("email = ?", email).First(&user).Error; errorUser != nil {
		return nil, errorUser
	}
	return &user, nil
}
