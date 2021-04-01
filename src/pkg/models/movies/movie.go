package movies

import(
	"example.com/amazingmovies/src/pkg/models"
)

type Movie struct {
	models.Model
	Title  string `gorm:"column:title;not null;" json:"title" form:"title"`
	Cast  []Star
	Director string `gorm:"column:director;not null;" json:"director" form:"director"`
	ReleaseYear int `gorm:"column:year;not null;" json:"year"`
	Poster string `gorm:"column:poster;" json:"poster"`
	Plot string `gorm:"column:plot;not null;" json:"plot"`
	Genres []Genre
  }

type Star struct{
	models.Model
	Name   string	`gorm:"column:name;not null;" json:"name"`
	BirthDate uint	`gorm:"column:birthdate;not null;" json:"birthdate"`
	BirthLocation uint `gorm:"column:birthlocation;not null;" json:"birthlocation"`
	// Maybe add the movies that starred in
}

type Genre struct{
	models.Model
	Type   string  `gorm:"column:type;not null;" json:"type"`
}

