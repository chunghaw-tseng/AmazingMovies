package persistence

import(
	"strings"
	"example.com/amazingmovies/src/pkg/db"
	models "example.com/amazingmovies/src/pkg/models/movies"
)

type GenreRepository struct{}
var genreRepository *GenreRepository

func GetGenreRepository() *GenreRepository {
	if genreRepository == nil {
		genreRepository = &GenreRepository{}
	}
	return genreRepository
}

func (r *GenreRepository) All() (*[]models.Genre, error) {
	var genre []models.Genre
	err := Find(&models.Genre{}, &genre, []string{}, "id asc")
	return &genre, err
}

func (r *GenreRepository) Query(q *models.Genre) (*[]models.Genre, error) {
	var genre []models.Genre
	err := Find(&q, &genre, []string{}, "id asc")
	return &genre, err
}

func (r *GenreRepository) GetFromType(type_genre string) (*models.Genre, error){
	var genre models.Genre
	where := models.Genre{}
	where.Type = strings.ToLower(type_genre)
	_, err := First(&where, &genre, []string{})
	if err != nil {
		return nil, err
	}
	return &genre, err
}

func (r *GenreRepository) Add(genre *models.Genre) (*models.Genre, error) {
	err := Create(&genre)
	err = Save(&genre)
	return genre, err
}

func (r *GenreRepository) Delete(genre *models.Genre) error { return db.GetDB().Unscoped().Delete(&genre).Error }