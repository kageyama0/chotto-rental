package auth_handler



type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

// @Description ユーザー登録リクエスト
type SignupRequest struct {
	DisplayName string `json:"displayName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
}

// @Description ログインリクエスト
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
