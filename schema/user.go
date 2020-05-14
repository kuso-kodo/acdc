package schema

type UserRegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
