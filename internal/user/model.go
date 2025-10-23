// Package user contains domain models and business logic for user management.
package user

// User represents a user entity in the chat application.
// It includes authentication credentials and basic profile information.
type User struct {
	ID       uint   `gorm:"primaryKey"`         // Primary key for database
	Name     string `json:"name"`               // User's display name
	Email    string `json:"email"`              // User's email address (unique)
	Password string `json:"password,omitempty"` // Bcrypt hashed password (omitted in responses)
}
