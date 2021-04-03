package persistence

import(
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

func (r *GenreRepository) Add(genre *models.Genre) error {
	err := Create(&genre)
	err = Save(&genre)
	return err
}

func (r *GenreRepository) Delete(genre *models.Genre) error { return db.GetDB().Unscoped().Delete(&genre).Error }