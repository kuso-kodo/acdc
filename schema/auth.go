package schema

type AuthLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Code   int    `json:"code"`
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
