package entity

import (
	"github.com/api_golang/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        entity.ID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        entity.NewID(),
		Name:      name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
