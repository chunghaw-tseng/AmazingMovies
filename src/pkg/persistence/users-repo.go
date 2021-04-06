package persistence

import(
	"example.com/amazingmovies/src/pkg/db"
	models "example.com/amazingmovies/src/pkg/models/users"
	"example.com/amazingmovies/src/pkg/models/movies"
	"strconv"
)

type UserRepository struct{}
var userRepository *UserRepository

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}


// Get user by api_key
func (r *UserRepository) GetbyKey(key string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.APIKey = key
	_, err := First(&where, &user, []string{"Favorites"})
	if err != nil {
		return nil, err
	}
	return &user, err
}


func (r *UserRepository) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{"Favorites"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{"Favorites"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Favorites"}, "id asc")
	return &users, err
}

func (r *UserRepository) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	return db.GetDB().Omit("RoleID").Save(&user).Error
}

func (r *UserRepository) Delete(user *models.User) error {
	db.GetDB().Model(user).Association("Favorites").Clear()
	return db.GetDB().Unscoped().Delete(&user).Error
}

func (r *UserRepository) DeleteAssociation(user *models.User ,association string, movie *movies.Movie) error {
	return db.GetDB().Model(user).Association(association).Delete(movie).Error
}