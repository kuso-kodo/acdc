package schema

type AuthLoginRequestSchema struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
