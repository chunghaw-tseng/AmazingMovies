package movies

import(
	"example.com/amazingmovies/src/pkg/models"
)

type People struct{
	models.BaseID
	Name   string	`gorm:"column:name;not null;" json:"name"`
	BirthDate string	`gorm:"column:birthdate;" json:"birthdate"`
	BirthLocation string `gorm:"column:birthlocation;" json:"birthlocation"`
	Gender string `gorm:"column:gender;" json:"gender"`

}

type Genre struct{
	models.BaseID
	Type   string  `gorm:"column:type;not null;" json:"type"`
}


type Movie struct {
	models.BaseID
	Title  string `gorm:"column:title;not null;" json:"title" form:"title"`
	Cast  []*People `gorm:"many2many:movie_cast;"`
	Director string `gorm:"column:director;not null;" json:"director" form:"director"`
	ReleaseYear string `gorm:"column:release_year;not null;" json:"release_year"`
	// Poster string `gorm:"column:poster;" json:"poster"`
	Plot string `gorm:"column:plot;not null;" json:"plot"`
	Genres []*Genre   `gorm:"many2many:movie_genres;"`
}