package user

import "go-chat-live/internal/database"

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindById(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}

type userRepositoryImpl struct{}

func NewUsuarioRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) Create(user *User) error {
	return database.DB.Create(user).Error
}

func (r *userRepositoryImpl) FindAll() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) FindById(id int) (*User, error) {
	var user User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Update(user *User) error {
	return database.DB.Save(user).Error
}

func (r *userRepositoryImpl) Delete(id int) error {
	return database.DB.Delete(&User{}, id).Error
}
