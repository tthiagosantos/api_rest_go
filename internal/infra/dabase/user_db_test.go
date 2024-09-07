package dabase

import (
	"github.com/api_golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUserCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John", "j@j.com", "123")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userNotFound entity.User
	err = db.First(&userNotFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userNotFound.ID)
	assert.Equal(t, user.Name, userNotFound.Name)
	assert.Equal(t, user.Email, userNotFound.Email)
	assert.NotNil(t, userNotFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John", "j@j.com", "123")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userNotFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userNotFound.ID)
	assert.Equal(t, user.Name, userNotFound.Name)
	assert.Equal(t, user.Email, userNotFound.Email)
	assert.NotNil(t, userNotFound.Password)
}
