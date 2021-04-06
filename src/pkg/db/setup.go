package db

import(
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
  		_ "github.com/jinzhu/gorm/dialects/sqlite"
	"example.com/amazingmovies/src/pkg/models/movies"
	"example.com/amazingmovies/src/pkg/models/users"
	"example.com/amazingmovies/src/pkg/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

func StartDatabase() {

	var database = DB
	var isInit bool

	configuration := config.GetConfig()

	// Used for SQL config
	driver := configuration.Database.Driver
	dbname := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "mysql"{
		isInit = checkAndcreateDB(username, password, host, port, dbname)
		dsn := username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local"
		database, err = gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println("db err: ", err)
			fmt.Println("Database not exists")
		}
	} 

	database.LogMode(false)
	database.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	database.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	database.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)	
	DB = database
	migration()
	// Init data for testing purposes
	if isInit {
		initRoles()
		initGenres()
		initPeople()
		initMovies()
	}
  }


  func checkAndcreateDB(username, password, host, port, dbname string) bool{
	var created = false
	create_db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/")
	if err != nil {
		panic(err)
	}
	defer create_db.Close()

	_,err = create_db.Exec("USE "+dbname)
	if err != nil {
		fmt.Println("Creating database ")
		// Create DB
		_,err = create_db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
		if err != nil {
			panic(err)
		}
		created = true
	}
	
	create_db.Close()
	return created
  }


  func migration(){
	DB.AutoMigrate(&movies.Movie{})
	DB.AutoMigrate(&movies.People{})
	DB.AutoMigrate(&movies.Genre{})
	DB.AutoMigrate(&users.UserRole{})
	DB.AutoMigrate(&users.User{})
  }

  func GetDB() *gorm.DB {
	return DB
}