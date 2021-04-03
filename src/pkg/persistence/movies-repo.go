package persistence

import (
	"fmt"
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

// TODO Fix these parts
func (r *MoviesRepository) Get(id string) (*models.Movie, error) {
	var movie models.Movie
	where := models.Movie{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	// The last item are associations i.e. other classes 
	_, err := First(&where, &movie, []string{"Genre"})
	if err != nil {
		return nil, err
	}
	return &movie, err
}

func (r *MoviesRepository) All() (*[]models.Movie, error) {
	var movies []models.Movie
	err := Find(&models.Movie{}, &movies, []string{"Genre"}, "id asc")
	return &movies, err
}

func (r *MoviesRepository) Query(q *models.Movie) (*[]models.Movie, error) {
	var movies []models.Movie
	err := Find(&q, &movies, []string{"Name"}, "id asc")
	return &movies, err
}

// TODO Find the Genre and add
func (r *MoviesRepository) Add(movies *models.Movie) error {
	fmt.Println(movies)
	err := Create(&movies)
	err = Save(&movies)
	return err
}

func (r *MoviesRepository) Update(movies *models.Movie) error { return db.GetDB().Omit("User").Save(&movies).Error }

func (r *MoviesRepository) Delete(movies *models.Movie) error { return db.GetDB().Unscoped().Delete(&movies).Error }