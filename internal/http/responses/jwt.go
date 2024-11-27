package responses

type JWTResponse struct {
	User        UserResponse
	AccessToken string
	ExpiresAt   int64
	ISS         string
}
