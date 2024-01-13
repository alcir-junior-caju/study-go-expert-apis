package database

import (
	"testing"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserDatabaseCreate(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.UserStruct{})
	user, _ := entity.User("John Doe", "johndoe@email.com", "password")
	userDatabase := UserDatabase(database)
	errorCreate := userDatabase.Create(user)
	assert.Nil(test, errorCreate)
	var userFound entity.UserStruct
	errorFound := database.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(test, errorFound)
	assert.Equal(test, user.ID, userFound.ID)
	assert.Equal(test, user.Name, userFound.Name)
	assert.Equal(test, user.Email, userFound.Email)
	assert.NotNil(test, userFound.Password)
}

func TestUserDatabaseFindByEmail(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.UserStruct{})
	user, _ := entity.User("John Doe", "johndoe@email.com", "password")
	userDatabase := UserDatabase(database)
	errorCreate := userDatabase.Create(user)
	assert.Nil(test, errorCreate)
	userFound, errorFound := userDatabase.FindByEmail(user.Email)
	assert.Nil(test, errorFound)
	assert.Equal(test, user.ID, userFound.ID)
	assert.Equal(test, user.Name, userFound.Name)
	assert.Equal(test, user.Email, userFound.Email)
	assert.NotNil(test, userFound.Password)

}
