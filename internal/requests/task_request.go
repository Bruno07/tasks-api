package requests

type TaskRequestDTO struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	User        UserRequestDTO `json:"user"`
}
