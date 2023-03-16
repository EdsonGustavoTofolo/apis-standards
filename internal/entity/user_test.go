package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@mail.com", "12345")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@mail.com", "12345")

	assert.NoError(t, err)

	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "123456", user.Password)
}
