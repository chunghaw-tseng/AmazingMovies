package db

import(
	"time"
	"github.com/jinzhu/gorm"
  		_ "github.com/jinzhu/gorm/dialects/sqlite"
	"example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/models/users"
	"example.com/amazingmovies/src/pkg/config"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

func StartDatabase() {

	var db = DB

	configuration := config.GetConfig()

	// Used for SQL config
	// driver := configuration.Database.Driver
	// database := configuration.Database.Dbname
	// username := configuration.Database.Username
	// password := configuration.Database.Password
	// host := configuration.Database.Host
	// port := configuration.Database.Port

	// TODO Choose what database to use maybe SQL
	// Used Sqlite for it to be easier
	// TODO Add Username and password 
	db, err := gorm.Open("sqlite3", "./data/amazingMovies.db")
  
	// TODO Get configuration
	if err != nil {
	  panic("Failed to connect to database!")
	}

	db.LogMode(false)
	db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()

  }

  func migration(){
	DB.AutoMigrate(&movies.Movie{})
	DB.AutoMigrate(&movies.People{})
	DB.AutoMigrate(&movies.Genre{})
	DB.AutoMigrate(&users.User{})
  }

  func GetDB() *gorm.DB {
	return DB
}