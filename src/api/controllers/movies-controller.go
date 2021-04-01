package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/db"
)

func GetMovies(c *gin.Context) {
	var movies []movies.Movie
	db.DB.Find(&movies)
  
	c.JSON(http.StatusOK, gin.H{"data": movies})
  }
  

func FindMovies(c *gin.Context) {  // Get model if exist
	var movie movies.Movie
  
	if err := db.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
  }