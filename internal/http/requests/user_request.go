package requests

type UserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ProfileId int64  `json:"profile_id"`
}
