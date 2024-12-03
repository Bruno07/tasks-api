package requests

type TaskRequestDTO struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	User        UserRequestDTO `json:"user"`
}
