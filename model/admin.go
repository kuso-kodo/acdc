package model

type Admin struct {
	UserID   uint   `gorm:"primary_key" json:"user_id"`
	UserName string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Role     uint   `gorm:"default:1"`
}

const (
	CustomerMask uint = 1 << iota
	ReceptionistMask
	MaintainerMask
	AccountingMask
	SuperUserMask
	InvalidMask
)

func (admin Admin) GetUserRole() []string {
	var result []string
	if admin.Role&CustomerMask != 0 {
		result = append(result, "customer")
	}
	if admin.Role&ReceptionistMask != 0 {
		result = append(result, "receptionist")
	}
	if admin.Role&MaintainerMask != 0 {
		result = append(result, "maintainer")
	}
	if admin.Role&AccountingMask != 0 {
		result = append(result, "accounting")
	}
	if admin.Role&SuperUserMask != 0 {
		result = append(result, "superuser")
	}
	return result
}

func (admin *Admin) IsCustomer() bool {
	return admin.Role&CustomerMask != 0
}

func (admin *Admin) IsReceptionist() bool {
	return admin.Role&ReceptionistMask != 0
}

func (admin *Admin) IsMaintainer() bool {
	return admin.Role&MaintainerMask != 0
}

func (admin *Admin) IsAccounting() bool {
	return admin.Role&AccountingMask != 0
}

func (admin *Admin) IsSuperUser() bool {
	return admin.Role&SuperUserMask != 0
}

func (admin *Admin) SetAsCustomer() {
	admin.Role |= CustomerMask
}

func (admin *Admin) SetAsReceptionist() {
	admin.Role |= ReceptionistMask
}

func (admin *Admin) SetAsMaintainer() {
	admin.Role |= MaintainerMask
}

func (admin *Admin) SetAsAccounting() {
	admin.Role |= AccountingMask
}

func (admin *Admin) SetAsSuperUser() {
	admin.Role |= SuperUserMask
}
