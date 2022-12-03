package request

type LoginReq struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	//Captcha  string `json:"captcha"`  // 验证码
}

type CreateAdminUserReq struct {
	Username string
	Password string
	Nickname string
	Email    string
	Avatar   string
}
