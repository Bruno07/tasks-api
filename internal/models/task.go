package models

import "errors"

type Task struct {
	ID          int64
	Title       string
	Description string
	UserID      int64
}

func (t *Task) Validate() error {

	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	if len(t.Description) > 2500 {
		return errors.New("Number of characters exceeded")
	}

	if t.UserID == 0 {
		return errors.New("User not found")
	}

	return nil
	
}
