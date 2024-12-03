package requests

type UserRequestDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ProfileID int64  `json:"profile_id"`
}
