package users
import(
	"example.com/amazingmovies/src/pkg/models"
)

type UserRole struct {
	models.BaseID
	RoleName string `gorm:"column:role_name;not null;" json:"role_name"`
	Users	[]User `gorm:"ForeignKey:RoleID"`
}
