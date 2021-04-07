package controllers

import(
	"log"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	models "example.com/amazingmovies/src/pkg/models/movies"
	persist "example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/http_errors"

)

type MovieInput struct {
	Title	string `json:"title" binding:"required"`
	Cast  	[]string `json:"cast" binding:"required"`
	Director string `json:"director" binding:"required"`
	ReleaseYear  string `json:"release_year" binding:"required"`
	Plot 	string `json:"plot" binding:"required"`
	Genre 	[]string `json:genre binding:"required"`
}

type MovieBasicOutput struct {
	ID 	uint64		`json:"id"`
	Title string	`json:"title"`
	ReleaseYear string	`json:"release_year"`
	Director string `json:"director"`
}

type PeopleInput struct{
	Name   string	`json:"name" binding:"required"`
	BirthDate string	`json:"birthdate" binding:"required"`
	BirthLocation string `json:"birthlocation" binding:"required"`
	Gender string `json:"gender" binding:"required"`
}


// GetMovies godoc
// @Summary Get all the movies or the ones specified by query
// @Description Get movies in BD
// @Produce json
// @Param query search string
// @Success 200 {object} []model.Movies
// @Router /am_api/movies [get]
func GetMovies(c *gin.Context) {
	s := persist.GetMovieRepository()
	var q models.Movie
	_ = c.Bind(&q)
	search := c.Query("search")
	if search == ""{
		fmt.Println("Search with no query")
		if movie, err := s.SimpleQuery(&q); err != nil {
			http_err.NewError(c, http.StatusNotFound, errors.New("movies not found"))
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, movie)
		}
	}else{
		if movie, err := s.QueryLike("title like ?", search+"%"); err != nil {
			http_err.NewError(c, http.StatusNotFound, errors.New("movies not found"))
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, movie)
		}
	}
  }


// CreateMovie godoc
// @Summary Creates a new entry for Movie
// @Description Adds a new movie entry in DB
// @Accept json
// @Param  title string
// @Param cast []string
// @Param director string
// @Param release_year string
// @Param plot string
// @Param genre []string
// @Produce json
// @Success 200 {object} []model.Movies
// @Router /am_api/movies [post]
func CreateMovie(c *gin.Context) {
	movie_repo := persist.GetMovieRepository()
	
	var movieInput MovieInput
	_ = c.BindJSON(&movieInput)

	// Create a new 
	new_movie := models.Movie{
		Title: movieInput.Title,
		Cast : getAndCreatePeople(movieInput.Cast),
		Director : movieInput.Director,
		ReleaseYear : movieInput.ReleaseYear,
		Plot : movieInput.Plot,
		Genres: getAndCreateGenres(movieInput.Genre),
	}
	if err := movie_repo.Add(&new_movie); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, new_movie)
		}
}

// GetMoviesByID godoc
// @Summary Show Movie Details
// @Description Get movie details by ID
// @Param id integer "Movie.ID"
// @Produce json
// @Success 200 {object} model.Movie
// @Router /am_api/movies/{id} [get]
func GetMoviesById(c *gin.Context) { 
	s := persist.GetMovieRepository()
	id := c.Param("id")
	if movie, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("movie not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, movie)
	}

  }


// UpdateMovie godoc
// @Summary Update Movie Info
// @Description Update Movie from ID
// @Param  id integer "Movie.ID"
// @Accept json
// @Param  title string
// @Param cast []string
// @Param director string
// @Param release_year string
// @Param plot string
// @Param genre []string
// @Produce json
// @Success 200 {object} model.Movie
// @Router /am_api/movies/{id} [put]
  func UpdateMovie(c *gin.Context){
	s := persist.GetMovieRepository()
	id := c.Param("id")
	var movieInput MovieInput
	_ = c.BindJSON(&movieInput)
	if movie, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("No Movie found"))
		log.Println(err)
	} else {
		s.DeleteAssociations(movie, []string{"Cast","Genres"})
		movie.Title = movieInput.Title
		movie.Director = movieInput.Director
		movie.ReleaseYear = movieInput.ReleaseYear
		movie.Plot = movieInput.Plot
		movie.Cast = getAndCreatePeople(movieInput.Cast)
		movie.Genres = getAndCreateGenres(movieInput.Genre)
		if err := s.Update(movie); err != nil {
			http_err.NewError(c, http.StatusBadRequest, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusCreated, movie)
			}
	}
}


// DeleteMovie godoc
// @Summary Deletes Movie from DB
// @Description Delete movie from ID
// @Param  id integer "Movie.ID"
// @Success 200 
// @Router /am_api/movies/{id} [delete]
// @Security Authorization Token from Admin user
func DeleteMovie(c *gin.Context){
	s := persist.GetMovieRepository()
	id := c.Param("id")
	if movie, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("movie not found"))
		log.Println(err)
	} else {
		if err := s.Delete(movie); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}


func getAndCreatePeople(cast []string) ([]*models.People){
	var movie_cast []*models.People
	ppl_repo := persist.GetPeopleRepository()

	for _, v := range cast {
		// Find the cast and the genre
		found, err := ppl_repo.GetFromName(v)
		if err != nil {
			// Create 
			new_ppl := models.People{
				Name: v,
			}
			new , _ := ppl_repo.Add(&new_ppl)
			movie_cast = append(movie_cast, new)
		}else{
			movie_cast = append(movie_cast, found)
		}
	}
	return movie_cast

}


func getAndCreateGenres(genres []string) ([]*models.Genre){
	var movie_genres []*models.Genre
	genre_repo := persist.GetGenreRepository()

	for _, v := range genres {
		// Find the cast and the genre
		found, err := genre_repo.GetFromType(v)
		if err != nil {
			// Create 
			new_genre := models.Genre{
				Type: strings.ToLower(v),
			}
			new , _ := genre_repo.Add(&new_genre)
			movie_genres = append(movie_genres, new)
		}else{
			movie_genres = append(movie_genres, found)
		}
	}
	return movie_genres
}


// GetPeople godoc
// @Summary Get all the famous people starred in the movies DB
// @Description return all the people in DB
// @Produce json
// @Success 200 {object} []model.People
// @Router /am_api/people [get]
func GetPeople(c *gin.Context) {
	s := persist.GetPeopleRepository()
	var q models.People
	_ = c.Bind(&q)
	if people, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("No People found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, people)
	}
}

// UpdatePeople godoc
// @Summary Update People Info
// @Description Update People from ID
// @Accept json
// @Param  id integer "People.ID"
// @Produce json
// @Param  name string
// @Param birthdate string
// @Param birthlocation string
// @Param gender string
// @Success 200 {object} model.People
// @Router /am_api/people/{id} [put]
func UpdatePeople(c *gin.Context) {
	s := persist.GetPeopleRepository()
	id := c.Param("id")
	var peopleInput PeopleInput
	_ = c.BindJSON(&peopleInput)
	if ppl, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("No People found"))
		log.Println(err)
	} else {
		ppl.Name = peopleInput.Name
		ppl.BirthDate = peopleInput.BirthDate
		ppl.BirthLocation = peopleInput.BirthLocation
		ppl.Gender = peopleInput.Gender
		if err := s.Update(ppl); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, ppl)
		}
	}
}

// GetGenres godoc
// @Summary Get all the genres from the DB
// @Description Get genres in DB
// @Produce json
// @Success 200 {object} []model.Genres
// @Router /am_api/genres [get]
func GetGenres(c *gin.Context) {
	s := persist.GetGenreRepository()
	var q models.Genre
	_ = c.Bind(&q)
	if genre, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Genres not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, genre)
	}
  }
