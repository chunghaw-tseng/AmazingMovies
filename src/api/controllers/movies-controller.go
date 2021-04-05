package controllers

import(
	"log"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	models "example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/http_errors"

)

type MovieInput struct {
	Title	string `json:"username" binding:"required"`
	Cast  	string `json:"lastname"`
	Director string `json:"firstname"`
	ReleaseYear  string `json:"password" binding:"required"`
	Poster      string `json:"role"`
	Plot 	string `json:"plot"`
	Genres 	string `json:genres`
}


func GetMovies(c *gin.Context) {
	s := persistence.GetMovieRepository()
	var q models.Movie
	_ = c.Bind(&q)

	if movie, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("movies not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, movie)
	}
  }


func CreateMovie(c *gin.Context) {
	s := persistence.GetMovieRepository()
	var movieInput models.Movie
	_ = c.BindJSON(&movieInput)
	if err := s.Add(&movieInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, movieInput)
	}
}
  

func GetMoviesById(c *gin.Context) {  // Get model if exist
	s := persistence.GetMovieRepository()
	id := c.Param("id")
	if movie, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("movie not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, movie)
	}

  }


  func GetGenres(c *gin.Context) {
	s := persistence.GetGenreRepository()
	var q models.Genre
	_ = c.Bind(&q)

	if genre, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Genres not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, genre)
	}
  }

  func CreateGenre(c *gin.Context) {
	s := persistence.GetGenreRepository()
	var genreInput models.Genre
	_ = c.BindJSON(&genreInput)
	if err := s.Add(&genreInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, genreInput)
	}
}


// TODO Delete genre by id
// func DeleteGenre(c *gin.Context) {
// 	s := persistence.GetGenreRepository()
// 	id := c.Params.ByName("id")
// 	var genreInput models.Genre
// 	_ = c.BindJSON(&genreInput)
// 	if genre, err := s.Get(id); err != nil {
// 		http_err.NewError(c, http.StatusNotFound, errors.New("Genre not found"))
// 		log.Println(err)
// 	} else {
// 		if err := s.Delete(genre); err != nil {
// 			http_err.NewError(c, http.StatusNotFound, err)
// 			log.Println(err)
// 		} else {
// 			c.JSON(http.StatusNoContent, "")
// 		}
// 	}
// }