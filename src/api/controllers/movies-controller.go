package controllers

import(
	"log"
	"errors"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	models "example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/http_errors"

)

type MovieInput struct {
	Title	string `json:"title" binding:"required"`
	Cast  	[]string `json:"cast"`
	Director string `json:"director"`
	ReleaseYear  string `json:"release_year"`
	Poster      string `json:"poster"`
	Plot 	string `json:"plot"`
	Genre 	[]string `json:genre`
}

// type SearchInput struct{
// 	Search string `json:"string" binding:"required"`
// } 


// Search
func GetMovies(c *gin.Context) {
	s := persistence.GetMovieRepository()
	var q models.Movie
	_ = c.Bind(&q)
	search := c.Query("search")
	if search == ""{
		if movie, err := s.Query(&q); err != nil {
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


// Find the genre and the cast
func CreateMovie(c *gin.Context) {
	movie_repo := persistence.GetMovieRepository()
	
	var movieInput MovieInput
	_ = c.BindJSON(&movieInput)

	// Create a new 
	new_movie := models.Movie{
		Title: movieInput.Title,
		Cast : getAndCreatePeople(movieInput.Cast),
		Director : movieInput.Director,
		ReleaseYear : movieInput.ReleaseYear,
		Poster	: movieInput.Poster,
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


func getAndCreatePeople(cast []string) ([]*models.People){
	var movie_cast []*models.People
	ppl_repo := persistence.GetPeopleRepository()

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
	genre_repo := persistence.GetGenreRepository()

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