package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		email         string
		password      string
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid user",
			username:    "testuser",
			email:       "test@example.com",
			password:    "password123",
			expectError: false,
		},
		{
			name:          "Empty username",
			username:      "",
			email:         "test@example.com",
			password:      "password123",
			expectError:   true,
			errorContains: "Username",
		},
		{
			name:          "Invalid email",
			username:      "testuser",
			email:         "invalid-email",
			password:      "password123",
			expectError:   true,
			errorContains: "Email",
		},
		{
			name:          "Password too short",
			username:      "testuser",
			email:         "test@example.com",
			password:      "pass",
			expectError:   true,
			errorContains: "Password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.username, tt.email, tt.password)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.username, user.Username)
				assert.Equal(t, tt.email, user.Email)
				assert.Empty(t, user.Password)
				assert.NotEmpty(t, user.HashPassword)
				assert.NotEmpty(t, user.ID)
				assert.Empty(t, user.Todos)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	originalPassword := "testpassword123"

	t.Run("Correct password comparison", func(t *testing.T) {
		user, err := NewUser("testuser", "test@example.com", originalPassword)
		assert.NoError(t, err)

		err = user.ComparePassword(originalPassword)
		assert.NoError(t, err)
	})

	t.Run("Incorrect password comparison", func(t *testing.T) {
		user, err := NewUser("testuser", "test@example.com", originalPassword)
		assert.NoError(t, err)

		err = user.ComparePassword("wrongpassword")
		assert.Error(t, err)
		assert.Equal(t, bcrypt.ErrMismatchedHashAndPassword, err)
	})
}

func TestAddHashPassword(t *testing.T) {
	t.Run("Successfully hash password", func(t *testing.T) {
		user := &User{
			ID:       "test-id",
			Username: "testuser",
			Email:    "test@example.com",
			Password: "testpassword123",
		}

		err := addHashPassword(user)
		assert.NoError(t, err)
		assert.Empty(t, user.Password)
		assert.NotEmpty(t, user.HashPassword)

		// Verify the hash works
		err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte("testpassword123"))
		assert.NoError(t, err)
	})
}
