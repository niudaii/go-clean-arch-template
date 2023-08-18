package entity

type LoginRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captchaId"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginUsecase interface {
	GetUserByUsername(username string) (user User, err error)
	CreateAccessToken(user *User) (accessToken string, err error)
}
