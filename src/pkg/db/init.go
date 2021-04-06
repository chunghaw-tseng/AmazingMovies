package db

import(
	"fmt"
	"strings"
	"github.com/google/uuid"
	"example.com/amazingmovies/pkg/crypto"
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
	// Create Admin
	var admin = users.User{
		Username : "admin",
		Firstname : "Amazing",
		Lastname  : "Movies",
		Hash      : crypto.GenerateHash([]byte("admin")),
		APIKey	  : uuid.New().String(),
		RoleID : 1,
	}
	db.Create(&admin)
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
		{"name":"Robert Downey Jr.", "birthdate":"04/04/1965", "birthLocation": "USA", "gender":"Male"},

	}
	for _, v := range data {
		var people = movies.People{Name: v["name"], BirthDate : v["birthdate"], BirthLocation: v["birthLocation"], Gender: v["gender"]}
		db.Create(&people)
	  }	
	fmt.Println("People init")
}

func initMovies(){
	db := GetDB()
	data := []map[string]string{
		{"title":"The Avengers", "director":"Joss Whedon", "releaseyear": "2012",
		"plot":"Earth's mightiest heroes must come together and learn to fight as a team if they are going to stop the mischievous Loki and his alien army from enslaving humanity.", 
		"cast": "Chris Evans",
		"genre":"Action"},
		
		{"title":"Wonder Woman", "director":"Patty Jenkins", "releaseyear": "2017", 
		"plot":"When a pilot crashes and tells of conflict in the outside world, Diana, an Amazonian warrior in training, leaves home to fight a war, discovering her full powers and true destiny.", 
		"cast": "Gal Gadot", 
		"genre": "Romance"},
		
		{"title":"Argo", "director":"Ben Affleck", "releaseyear": "2012", 
		"plot":"Acting under the cover of a Hollywood producer scouting a location for a science fiction film, a CIA agent launches a dangerous operation to rescue six Americans in Tehran during the U.S. hostage crisis in Iran in 1979.", 
		"cast": "Ben Affleck",
		"genre":"Thriller"},
		
		{"title":"Avengers Infinity War", "director":"Anthony Russo", "releaseyear": "2018", 
		"plot":"The Avengers and their allies must be willing to sacrifice all in an attempt to defeat the powerful Thanos before his blitz of devastation and ruin puts an end to the universe.", 
		"cast": "Elizabeth Olsen", 
		"genre":"Adventure"},
	}
	
	for _, m := range data{
		
		var ppl = movies.People {
			Name: m["cast"],
		} 
		db.Where(ppl).FirstOrInit(&ppl)
	
		var gn = movies.Genre {
			Type: strings.ToLower(m["genre"]),
		} 
		db.Where(gn).FirstOrInit(&gn)
			
		var movie = movies.Movie{
			Title: m["title"],
			Director: m["director"],
			ReleaseYear: m["releaseyear"],
			Plot: m["plot"],
			Cast: []*movies.People{&ppl},
			Genres:[]*movies.Genre{&gn},
		}
		db.Create(&movie)
	}
	fmt.Println("Movies init")
}