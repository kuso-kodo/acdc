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

func (user Admin) GetUserRole() []string {
	var result []string
	if user.Role&CustomerMask != 0 {
		result = append(result, "customer")
	}
	if user.Role&ReceptionistMask != 0 {
		result = append(result, "receptionist")
	}
	if user.Role&MaintainerMask != 0 {
		result = append(result, "maintainer")
	}
	if user.Role&AccountingMask != 0 {
		result = append(result, "accounting")
	}
	if user.Role&SuperUserMask != 0 {
		result = append(result, "superuser")
	}
	return result
}

func (user *Admin) IsCustomer() bool {
	return user.Role&CustomerMask != 0
}

func (user *Admin) IsReceptionist() bool {
	return user.Role&ReceptionistMask != 0
}

func (user *Admin) IsMaintainer() bool {
	return user.Role&MaintainerMask != 0
}

func (user *Admin) IsAccounting() bool {
	return user.Role&AccountingMask != 0
}

func (user *Admin) IsSuperUser() bool {
	return user.Role&SuperUserMask != 0
}

func (user *Admin) SetAsCustomer() {
	user.Role |= CustomerMask
}

func (user *Admin) SetAsReceptionist() {
	user.Role |= ReceptionistMask
}

func (user *Admin) SetAsMaintainer() {
	user.Role |= MaintainerMask
}

func (user *Admin) SetAsAccounting() {
	user.Role |= AccountingMask
}

func (user *Admin) SetAsSuperUser() {
	user.Role |= SuperUserMask
}
