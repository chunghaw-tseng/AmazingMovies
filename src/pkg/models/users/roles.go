package users

type UserRole struct {
	RoleID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	RoleName string `gorm:"column:role_name;not null;" json:"role_name"`
	Users	[]User `gorm:"ForeignKey:RoleID"`
}
