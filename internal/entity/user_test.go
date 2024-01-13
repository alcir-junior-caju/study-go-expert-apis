package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(test *testing.T) {
	user, error := User("John Doe", "johndoe@email.com", "password")
	assert.Nil(test, error)
	assert.NotNil(test, user)
	assert.NotEmpty(test, user.ID)
	assert.NotEmpty(test, user.Password)
	assert.Equal(test, "John Doe", user.Name)
	assert.Equal(test, "johndoe@email.com", user.Email)
}

func TestUserValidatePassword(test *testing.T) {
	user, error := User("John Doe", "johndoe@email.com", "password")
	assert.Nil(test, error)
	assert.True(test, user.ValidatePassword("password"))
	assert.False(test, user.ValidatePassword("wrong-password"))
	assert.NotEqual(test, "password", user.Password)
}
