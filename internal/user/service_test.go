package user

import (
	"errors"
	"testing"
)

// Mock do repository
type mockUserRepo struct {
	mockCreate      func(*User) error
	mockFindById    func(int) (*User, error)
	mockFindByEmail func(string) (*User, error)
}

func (m *mockUserRepo) Create(u *User) error {
	return m.mockCreate(u)
}
func (m *mockUserRepo) FindAll() ([]User, error) { return nil, nil }
func (m *mockUserRepo) FindById(id int) (*User, error) {
	return m.mockFindById(id)
}
func (m *mockUserRepo) FindByEmail(email string) (*User, error) {
	return m.mockFindByEmail(email)
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

	u := &User{Name: "Guilherme", Email: "gui@email.com", Password: "123456"}
	err := Create(u)

	if err != nil {
		t.Errorf("expected nil, but got error: %v", err)
	}
}

func TestCreateUser_NoName(t *testing.T) {
	mockRepo := &mockUserRepo{
		mockCreate: func(u *User) error {
			return nil // não importa, service deve falhar antes
		},
	}

	repo = mockRepo

	u := &User{Name: "", Email: "gui@email.com", Password: "123456"}
	err := Create(u)

	if err == nil {
		t.Error("esperava erro ao criar usuário sem nome, mas veio nil")
	}
}

func TestCreateUser_NoEmail(t *testing.T) {
	mockRepo := &mockUserRepo{
		mockCreate: func(u *User) error {
			return nil
		},
	}

	repo = mockRepo

	u := &User{Name: "Guilherme", Email: "", Password: "123456"}
	err := Create(u)

	if err == nil {
		t.Error("esperava erro ao criar usuário sem email, mas veio nil")
	}
}

func TestLogin_WithValidCredentials(t *testing.T) {
	// Hash da senha "123456" usando bcrypt
	hashedPassword := "$2a$10$lzNEdWrZLsC4V5jcUZ5rXOp0S6SPsKCaO040IJwn.KKSF8yEJlLIq"

	mockRepo := &mockUserRepo{
		mockFindByEmail: func(email string) (*User, error) {
			if email == "test@email.com" {
				return &User{
					ID:       1,
					Name:     "Test User",
					Email:    "test@email.com",
					Password: hashedPassword,
				}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	repo = mockRepo

	response, err := Login("test@email.com", "123456")

	if err != nil {
		t.Errorf("expected successful login, but got error: %v", err)
	}

	if response == nil {
		t.Error("expected login response, but got nil")
	}

	if response != nil && response.Token == "" {
		t.Error("expected token in response, but got empty string")
	}

	if response != nil && response.User.Name != "Test User" {
		t.Errorf("expected user name 'Test User', but got '%s'", response.User.Name)
	}
}

func TestLogin_WithInvalidCredentials(t *testing.T) {
	mockRepo := &mockUserRepo{
		mockFindByEmail: func(email string) (*User, error) {
			return nil, errors.New("user not found")
		},
	}

	repo = mockRepo

	response, err := Login("invalid@email.com", "wrongpassword")

	if err == nil {
		t.Error("expected error for invalid credentials, but got nil")
	}

	if response != nil {
		t.Error("expected nil response for invalid credentials, but got response")
	}
}

func TestLogin_WithWrongPassword(t *testing.T) {
	// Hash da senha "123456"
	hashedPassword := "$2a$10$lzNEdWrZLsC4V5jcUZ5rXOp0S6SPsKCaO040IJwn.KKSF8yEJlLIq"

	mockRepo := &mockUserRepo{
		mockFindByEmail: func(email string) (*User, error) {
			return &User{
				ID:       1,
				Name:     "Test User",
				Email:    "test@email.com",
				Password: hashedPassword,
			}, nil
		},
	}

	repo = mockRepo

	response, err := Login("test@email.com", "wrongpassword")

	if err == nil {
		t.Error("expected error for wrong password, but got nil")
	}

	if response != nil {
		t.Error("expected nil response for wrong password, but got response")
	}
}
