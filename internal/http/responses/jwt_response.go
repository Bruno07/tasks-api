package responses

type JWTResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
	ExpiresAt   int64        `json:"expires_at"`
	ISS         string       `json:"iss"`
}
