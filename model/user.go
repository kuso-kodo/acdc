package model

type User struct {
	UserID   uint   `gorm:"primary_key" json:"user_id"`
	UserName string `gorm:"default:'AC/DC'" json:"username"`
	Password string `json:"-"`
	Role     uint   `gorm:"default:1"`
	Phone    string `gorm:"unique" json:"phone"`
}

const (
	CustomerMask uint = 1 << iota
	ReceptionistMask
	MaintainerMask
	AccountingMask
	SuperUserMask
	InvalidMask
)

func (user User) GetUserRole() []string {
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

func (user *User) IsCustomer() bool {
	return user.Role&CustomerMask != 0
}

func (user *User) IsReceptionist() bool {
	return user.Role&ReceptionistMask != 0
}

func (user *User) IsMaintainer() bool {
	return user.Role&MaintainerMask != 0
}

func (user *User) IsAccounting() bool {
	return user.Role&AccountingMask != 0
}

func (user *User) IsSuperUser() bool {
	return user.Role&SuperUserMask != 0
}

func (user *User) SetAsCustomer() {
	user.Role |= CustomerMask
}

func (user *User) SetAsReceptionist() {
	user.Role |= ReceptionistMask
}

func (user *User) SetAsMaintainer() {
	user.Role |= MaintainerMask
}

func (user *User) SetAsAccounting() {
	user.Role |= AccountingMask
}

func (user *User) SetAsSuperUser() {
	user.Role |= SuperUserMask
}
