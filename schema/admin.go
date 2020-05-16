package schema

type AddNewAdminRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type AddNewAdminResponse CommonStatusSchema

type DeleteAdminRequest struct {
	UserName string `json:"username"`
}

type DeleteAdminResponse CommonStatusSchema
