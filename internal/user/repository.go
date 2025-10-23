package user

import "go-chat-live/internal/database"

// UserRepository defines the interface for user data access operations.
// This interface follows the Repository pattern to abstract database operations.
type UserRepository interface {
	Create(user *User) error                 // Creates a new user record
	FindAll() ([]User, error)                // Retrieves all users
	FindById(id int) (*User, error)          // Finds user by ID
	FindByEmail(email string) (*User, error) // Finds user by email address
	Update(user *User) error                 // Updates existing user
	Delete(id int) error                     // Deletes user by ID
}

// userRepositoryImpl implements UserRepository using GORM ORM.
type userRepositoryImpl struct{}

// NewUsuarioRepository creates a new instance of UserRepository.
func NewUsuarioRepository() UserRepository {
	return &userRepositoryImpl{}
}

// Create inserts a new user into the database.
func (r *userRepositoryImpl) Create(user *User) error {
	return database.DB.Create(user).Error
}

// FindAll retrieves all users from the database.
func (r *userRepositoryImpl) FindAll() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error
	return users, err
}

// FindById retrieves a user by their ID.
func (r *userRepositoryImpl) FindById(id int) (*User, error) {
	var user User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email address.
// Used primarily for authentication purposes.
func (r *userRepositoryImpl) FindByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update saves changes to an existing user record.
func (r *userRepositoryImpl) Update(user *User) error {
	return database.DB.Save(user).Error
}

// Delete removes a user record by ID.
func (r *userRepositoryImpl) Delete(id int) error {
	return database.DB.Delete(&User{}, id).Error
}
