package db

import(
	"github.com/jinzhu/gorm"
  		_ "github.com/jinzhu/gorm/dialects/sqlite"
	"example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/models/users"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Used Sqlite for it to be easier
	// TODO Add Username and password 
	database, err := gorm.Open("sqlite3", "./data/amazingMovies.db")
  

	// TODO Get configuration

	if err != nil {
	  panic("Failed to connect to database!")
	}
  
	DB = database
	migration()

  }

  func migration(){
	DB.AutoMigrate(&movies.Movie{})
	DB.AutoMigrate(&movies.Star{})
	DB.AutoMigrate(&movies.Genre{})
	DB.AutoMigrate(&users.User{})
  }

  func GetDB() *gorm.DB {
	return DB
}