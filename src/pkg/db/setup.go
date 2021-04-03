package db

import(
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
  		_ "github.com/jinzhu/gorm/dialects/sqlite"
		_ "github.com/jinzhu/gorm/dialects/mysql"
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

	configuration := config.GetConfig()

	// Used for SQL config
	driver := configuration.Database.Driver
	dbname := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "mysql"{
		createDB(username, password, host, port, dbname)
		dsn := username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local"
		database, err = gorm.Open("mysql", dsn)
		if err != nil {
			panic("Failed to connect to database!")
			fmt.Println("db err: ", err)
		}
	} 
	// else if driver == "sqlite"{
	// 		dsn := "./data/"+dbname+".db"
	// 		database, err := gorm.Open("sqlite3", dsn)
	// 		if err != nil {
	// 			panic("Failed to connect to database!")
	// 			fmt.Println("db err: ", err)
	// 		}
	// }

	database.LogMode(false)
	database.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	database.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	database.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = database
	migration()

  }

  func createDB(username, password, host, port, dbname string){
		create_db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/")
		if err != nil {
			panic(err)
		}
		defer create_db.Close()

		_,err = create_db.Exec("CREATE DATABASE IF NOT EXISTS "+dbname)
		if err != nil {
			panic(err)
		}
		fmt.Println("Database exists or it was created successfully")
		create_db.Close()
	 
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