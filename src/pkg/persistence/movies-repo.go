package persistence

import (
	"example.com/amazingmovies/src/pkg/db"
	models "example.com/amazingmovies/src/pkg/models/movies"
	"strconv"
)


type MoviesRepository struct{}
var movieRepository *MoviesRepository

func GetMovieRepository() *MoviesRepository {
	if movieRepository == nil {
		movieRepository = &MoviesRepository{}
	}
	return movieRepository
}

func (r *MoviesRepository) Get(id string) (*models.Movie, error) {
	var movie models.Movie
	where := models.Movie{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	// The last item are associations i.e. other classes 
	_, err := First(&where, &movie, []string{"Genres", "Cast"})
	if err != nil {
		return nil, err
	}
	return &movie, err
}

func (r *MoviesRepository) All() (*[]models.Movie, error) {
	var movies []models.Movie
	err := Find(&models.Movie{}, &movies, []string{"Genres", "Cast"}, "id asc")
	return &movies, err
}

func (r *MoviesRepository) Query(q *models.Movie) (*[]models.Movie, error) {
	var movies []models.Movie
	err := Find(&q, &movies, []string{"Genres", "Cast"}, "id asc")
	return &movies, err
}

func (r *MoviesRepository) QueryLike(column string, query string) (*[]models.Movie, error) {
	var movies []models.Movie
	err := FindLike(column, query, &movies, []string{"Genres", "Cast"}, "id asc")
	return &movies, err
}


func (r *MoviesRepository) Add(movies *models.Movie) error {
	err := Create(&movies)
	err = Save(&movies)
	return err
}

func (r *MoviesRepository) Update(movies *models.Movie) error { return db.GetDB().Save(&movies).Error }

func (r *MoviesRepository) Delete(movies *models.Movie) error { return db.GetDB().Unscoped().Delete(&movies).Error }