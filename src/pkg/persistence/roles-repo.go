package persistence

import(
	models "example.com/amazingmovies/src/pkg/models/users"
)

type RolesRepository struct{}
var rolesRepository *RolesRepository

func GetRolesRepository() *RolesRepository {
	if rolesRepository == nil {
		rolesRepository = &RolesRepository{}
	}
	return rolesRepository
}

func (r *RolesRepository) Get(rl string) (*models.UserRole, error) {
	var role models.UserRole
	where := models.UserRole{}
	where.RoleName = rl
	// The last item are associations i.e. other classes 
	_, err := First(&where, &role, []string{})
	if err != nil {
		return nil, err
	}
	return &role, err
}