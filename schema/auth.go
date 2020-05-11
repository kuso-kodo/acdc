package schema

type AuthLoginRequestSchema struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AuthLoginResponseSchema struct {
	Code   int    `json:"code"`
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
