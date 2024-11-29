package responses

import "time"

type TaskResponse struct {
	ID            int64        `json:"id,omitempty"`
	Summary       string       `json:"summary,omitempty"`
	PerformedDate *time.Time   `json:"performed_date,omitempty"`
	User          UserResponse `json:"user,omitempty"`
}
