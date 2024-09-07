package dabase

import (
	"github.com/api_golang/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Create(user *entity.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindByID(id string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update(user *entity.User) error {
	_, err := u.FindByID(user.ID.String())
	if err != nil {
		return err
	}
	if err := u.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(id string) error {
	_, err := u.FindByID(id)
	if err != nil {
		return err
	}
	if err := u.DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindAll(page, limit int, sort string) ([]entity.User, error) {
	var users []entity.User
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&users).Error
	} else {
		err = u.DB.Limit(limit).Order("created_at " + sort).Find(&users).Error
	}
	return users, err
}
