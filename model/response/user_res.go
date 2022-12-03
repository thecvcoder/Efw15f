package response

type UserResponse struct {
}

// LoginResponse 所有用户登录响应
type LoginResponse struct {
	// 暂时留空，应该返回UUID或者用户名，容我考虑一下
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
