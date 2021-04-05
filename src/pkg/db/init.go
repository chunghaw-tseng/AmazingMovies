package db

import(
	"fmt"
	"example.com/amazingmovies/src/pkg/models/users"
	"example.com/amazingmovies/src/pkg/models/movies"
)

func initRoles(){
	db := GetDB()
	data := []string{
		"admin",
		"user",
	 }
	 for _, v := range data {
		var role = users.UserRole{RoleName: v}
		db.Create(&role)
	  }	
	fmt.Println("Roles init")
}

func initGenres(){
	db := GetDB()
	data := []string{
		"Action",
		"Horror",
		"Adventure",
		"Comedy",
		"Sci-Fi",
		"Crime",
		"Thriller",
	 }
	 for _, v := range data {
		var genre = movies.Genre{Type: v}
		db.Create(&genre)
	  }	
	fmt.Println("Genre init")

}