package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "1234")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "1234", user.Password)
}
