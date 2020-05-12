package schema

type AddNewAdminRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type AddNewAdminResponse CommonFailureSchema

type DeleteAdminRequest struct {
	UserName string `json:"username"`
}

type DeleteAdminResponse CommonFailureSchema
