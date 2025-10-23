package user

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// repo is the global repository instance used by service functions
var repo = NewUsuarioRepository()

// getJWTSecret retrieves JWT secret from environment variables with fallback
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key"
	}
	return []byte(secret)
}

// Create validates and creates a new user with required fields validation
func Create(user *User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return repo.Create(user)
}

// List retrieves all users from the repository
func List() ([]User, error) {
	return repo.FindAll()
}

// FindById retrieves a specific user by ID with error handling for not found cases
func FindById(id int) (*User, error) {
	user, err := repo.FindById(id)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

// Update modifies an existing user's information
func Update(id int, newData *User) (*User, error) {
	user, err := repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Name = newData.Name
	user.Email = newData.Email

	err = repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete removes a user by ID with validation
func Delete(id int) error {
	user, err := FindById(id)
	if err != nil || user == nil {
		return fmt.Errorf("user not found")
	}
	return repo.Delete(id)
}

// LoginRequest represents the payload for user authentication
type LoginRequest struct {
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's plain text password
}

// LoginResponse contains authentication result with JWT token and user data
type LoginResponse struct {
	Token string `json:"token"` // JWT access token
	User  User   `json:"user"`  // User information (password omitted)
}

// Login authenticates user credentials and returns JWT token
// Validates email/password combination using bcrypt and generates JWT
func Login(email, password string) (*LoginResponse, error) {
	user, err := repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password hash using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Create JWT token with user claims and 24h expiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return nil, errors.New("error generating token")
	}

	user.Password = "" // Remove password from response for security
	return &LoginResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}
