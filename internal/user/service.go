package user

import (
	"errors"
	"fmt"
)

var repo = NewUsuarioRepository()

func Create(user *User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return repo.Create(user)
}

func Listar() ([]User, error) {
	return repo.FindAll()
}

func FindById(id int) (*User, error) {
	user, err := repo.FindById(id)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}

	return user, nil
}

func Update(id int, newData *User) (*User, error) {
	user, err := repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Atualizar os campos (pode ser validado campo a campo também)
	user.Name = newData.Name
	user.Email = newData.Email

	err = repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Delete(id int) error {
	user, err := FindById(id)
	if err != nil || user == nil {
		return fmt.Errorf("user not found")
	}

	return repo.Delete(id)
}
