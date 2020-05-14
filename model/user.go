package model

type User struct {
	UserID   uint     `gorm:"primary_key" json:"user_id"`
	UserName string   `gorm:"unique" json:"username"`
	Password string   `json:"-"`
	Phone    string   `gorm:"unique" json:"phone"`
	Tickets  []Ticket `json:"-" gorm:"foreignkey:UserRefer;association_foreignkey:UserID"`
}
