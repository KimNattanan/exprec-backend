package dto

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// type LoginUserResponse struct {
// 	SessionID             string    `json:"session_id"`
// 	AccessToken           string    `json:"access_token"`
// 	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
// 	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
// 	RefreshToken          string    `json:"refresh_token"`
// 	User *UserResponse `json:"user"`
// }
