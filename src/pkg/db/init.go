package db

import(
	"fmt"
	"strings"
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

// func initAdmin(){
// 	db := GetDB()
// 	var admin = users.User{RoleName: v}
// 	db.Create(&role)
// }


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
		var genre = movies.Genre{Type: strings.ToLower(v)}
		db.Create(&genre)
	  }	
	fmt.Println("Genre init")
}

func initPeople(){
	db := GetDB()
	data := []map[string]string{
		{"name":"Gal Gadot", "birthdate":"30/02/1985", "birthLocation": "Israel", "gender":"Female"},
		{"name":"Elizabeth Olsen", "birthdate":"16/02/1989", "birthLocation": "USA", "gender":"Female"},
		{"name":"Ben Affleck", "birthdate":"15/08/1972", "birthLocation": "USA", "gender":"Male"},
		{"name":"Chris Evans", "birthdate":"13/06/1981", "birthLocation": "USA", "gender":"Male"},
	}
	for _, v := range data {
		var people = movies.People{Name: v["name"], BirthDate : v["birthdate"], BirthLocation: v["birthLocation"], Gender: v["gender"]}
		db.Create(&people)
	  }	
	fmt.Println("People init")
}

// func initMovies(){
// 	db := GetDB()

// }