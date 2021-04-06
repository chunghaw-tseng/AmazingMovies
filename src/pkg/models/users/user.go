package users

import(
	"example.com/amazingmovies/src/pkg/models"
	"example.com/amazingmovies/src/pkg/models/movies"
	"time"
)

type User struct{
	models.Base
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string   `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	// Hash for password
	Hash      string   `gorm:"column:hash;not null;" json:"hash"`
	APIKey	  string   `gorm:"column:apikey;not null;" json:"apikey"`
 	RoleID      uint64 
	// Role 		UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Favorites []*movies.Movie  `gorm:"many2many:userfav_movies;"`
}

func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}