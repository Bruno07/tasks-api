package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests whether the fields are filled out correctly
func TestUserModel_Validate(t *testing.T) {

	var user = User{
		Name:     "Test 1",
		Email:    "test1@email.com",
		Password: "12345678",
	}

	t.Run("Should return no validation errors", func(t *testing.T) {

		err := user.Validate()

		assert.NoError(t, err)

	})

	t.Run("Should return an empty name validation error", func(t *testing.T) {

		user.Name = ""
		err := user.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "Name is required")
	})

	t.Run("Should return an empty email validation error", func(t *testing.T) {

		user.Name = "Test 1"
		user.Email = ""
		err := user.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "Email is required")
	})

	t.Run("Must return a invalid email", func(t *testing.T) {

		user.Name = "Test 1"
		user.Email = "test1email.com"
		user.Password = "12345678"
		err := user.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "Invalid email format")
	})

	t.Run("Must return an invalid password", func(t *testing.T) {

		user.Name = "Test 1"
		user.Email = "test1@email.com"
		user.Password = "123456"
		err := user.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "Password must be at least 8 characters long")
	})

}

// Tests whether the fields are filled out correctly
func TestUserModel_HashPassword(t *testing.T) {

	var user = User{Password: "12345678"}
	err := user.HashPassword()

	assert.NoError(t, err)

}

// Check if the entered password exists
func TestUserModel_CheckPassword(t *testing.T) {

	var user = User{Password: "$2a$10$tAN6tScsHD4hcdNhYlySgO28hH1rv9v2cOPqt6j5ebyokGk7UG9Wi"}
	isValid := user.CheckPassword("12345678")

	assert.True(t, isValid)

}

func TestUserModel_GetPermissions(t *testing.T) {

	t.Run("Must return permissions with technician profile", func(t *testing.T) {
		
		var user = User{ProfileID: Technical}
		permissions := user.GetPermissions()
	
		assert.Equal(t, 3, len(permissions))
		assert.Equal(t, permissions[0], "tasks:create")
		assert.Equal(t, permissions[1], "tasks:update")
		assert.Equal(t, permissions[2], "tasks:view")

	})

	t.Run("Must return permissions with manager profile", func(t *testing.T) {
		
		var user = User{ProfileID: Manager}
		permissions := user.GetPermissions()
	
		assert.Equal(t, 4, len(permissions))
		assert.Equal(t, permissions[0], "tasks:create")
		assert.Equal(t, permissions[1], "tasks:update")
		assert.Equal(t, permissions[2], "tasks:view")
		assert.Equal(t, permissions[3], "tasks:delete")

	})


}

