package persistence

import(
	"strings"
	"strconv"
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

func (r *GenreRepository) Get(id string) (*models.Genre, error) {
	var genre models.Genre
	where := models.Genre{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &genre, []string{})
	if err != nil {
		return nil, err
	}
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