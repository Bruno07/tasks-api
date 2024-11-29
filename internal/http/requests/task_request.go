package requests

import "time"

type TaskRequest struct {
	ID            int64
	Summary       string `json:"summary"`
	PerformedDate time.Time
	UserID        int64
}
