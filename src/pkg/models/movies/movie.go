package movies

import(
	"example.com/amazingmovies/src/pkg/models"
)

type People struct{
	models.Base
	Name   string	`gorm:"column:name;not null;" json:"name"`
	BirthDate string	`gorm:"column:birthdate;not null;" json:"birthdate"`
	BirthLocation string `gorm:"column:birthlocation;not null;" json:"birthlocation"`
	DeathDate string  `gorm:"column:deathdate;not null;" json:"deathdate"`
	Gender uint `gorm:"column:gender;not null;" json:"gender"`
}

type Genre struct{
	models.Base
	Type   string  `gorm:"column:type;not null;" json:"type"`
}


type Movie struct {
	models.Base
	Title  string `gorm:"column:title;not null;" json:"title" form:"title"`
	Cast  []People `gorm:"many2many:movie_cast;" json:"cast"`
	// Cast  string `gorm:"column:cast;not null;" json:"cast" form:"cast"`
	Director string `gorm:"column:director;not null;" json:"director" form:"director"`
	ReleaseYear int `gorm:"column:year;not null;" json:"year"`
	Poster string `gorm:"column:poster;" json:"poster"`
	Plot string `gorm:"column:plot;not null;" json:"plot"`
	Genres []Genre   `gorm:"many2many:movie_genres;" json:"genres"`
  }



