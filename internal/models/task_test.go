package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskModel_Validate(t *testing.T) {

	var task = Task{
		Title:       "Test Create Task",
		Description: "This is my creation test",
		UserID:      1,
	}

	t.Run("Should return no validation errors", func(t *testing.T) {

		err := task.Validate()

		assert.NoError(t, err)

	})

	t.Run("Should return an empty title validation error", func(t *testing.T) {

		task.Title = ""
		err := task.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "title is required")
	})

	t.Run("Should return an empty description validation error", func(t *testing.T) {

		task.Title = "Test Create Task"
		task.Description = ""
		err := task.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "description is required")
	})

	t.Run("Should return an empty user id validation error", func(t *testing.T) {

		task.Title = "Test Create Task"
		task.Description = "This is my creation test"
		task.UserID = 0
		err := task.Validate()

		assert.Error(t, err)
		assert.EqualError(t, err, "User not found")
	})

}
