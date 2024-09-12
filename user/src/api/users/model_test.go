package users

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser_ValidateEmail(t *testing.T) {
	tests := []struct {
		email   string
		isValid bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"another@test.org", true},
		{"invalid@", false},
	}

	for _, tt := range tests {
		user := User{Email: tt.email}
		if valid := user.ValidateEmail(); valid != tt.isValid {
			t.Errorf("expected email validation for %v to be %v, got %v", tt.email, tt.isValid, valid)
		}
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		user     User
		hasError bool
	}{
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "john@example.com", Age: 25}, false},
		{User{ID: primitive.NewObjectID(), Name: "", Email: "john@example.com", Age: 25}, true},      // Invalid name
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "invalid-email", Age: 25}, true},     // Invalid email
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "john@example.com", Age: -1}, true},  // Invalid age
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "john@example.com", Age: 0}, true},   // Invalid age
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "", Age: 25}, true},                  // Empty email
		{User{ID: primitive.NewObjectID(), Name: "John", Email: "john@example.com", Age: 25}, false}, // Valid user
	}

	for _, tt := range tests {
		err := tt.user.Validate()
		if (err != nil) != tt.hasError {
			t.Errorf("expected validation error: %v, got: %v", tt.hasError, err)
		}
	}
}
