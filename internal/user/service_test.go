package user

import (
	"errors"
	"testing"
)

// Mock do repository
type mockUserRepo struct {
	mockCreate   func(*User) error
	mockFindById func(int) (*User, error)
}

func (m *mockUserRepo) Create(u *User) error {
	return m.mockCreate(u)
}
func (m *mockUserRepo) FindAll() ([]User, error) { return nil, nil }
func (m *mockUserRepo) FindById(id int) (*User, error) {
	return m.mockFindById(id)
}
func (m *mockUserRepo) Update(u *User) error { return nil }
func (m *mockUserRepo) Delete(id int) error  { return nil }

func TestCreateUser_WithValidData(t *testing.T) {
	mockRepo := &mockUserRepo{
		mockCreate: func(u *User) error {
			if u.Name == "" || u.Email == "" {
				return errors.New("invalid data")
			}
			return nil
		},
	}

	repo = mockRepo // injetando mock no service

	u := &User{Name: "Guilherme", Email: "gui@email.com"}
	err := Create(u)

	if err != nil {
		t.Errorf("waits nil, but return error: %v", err)
	}
}

func TestCreateUser_NoName(t *testing.T) {
	mockRepo := &mockUserRepo{
		mockCreate: func(u *User) error {
			return nil // não importa, service deve falhar antes
		},
	}

	repo = mockRepo

	u := &User{Name: "", Email: "gui@email.com"}
	err := Create(u)

	if err == nil {
		t.Error("esperava erro ao criar usuário sem nome, mas veio nil")
	}
}
