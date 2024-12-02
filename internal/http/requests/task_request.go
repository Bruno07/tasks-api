package requests

import "time"

type TaskRequest struct {
	ID            int64     `json:"id"`
	Summary       string    `json:"summary"`
	PerformedDate time.Time `json:"performed_date"`
	UserID        int64     `json:"user_id"`
}
