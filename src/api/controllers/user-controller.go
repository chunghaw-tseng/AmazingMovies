package controllers

import(
	"log"
	"fmt"
	"net/http"
	"errors"
	models "example.com/amazingmovies/src/pkg/models/users"
	"example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/crypto"
	"example.com/amazingmovies/pkg/http_errors"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)


type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}


func GetUserById(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	s := persistence.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	r := persistence.GetRolesRepository()	
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	apikey := uuid.New().String()
	if userInput.Role == ""{
		userInput.Role = "user"
	}
	role, _ := r.Get(userInput.Role)

	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.GenerateHash([]byte(userInput.Password)),
		APIKey:	   apikey,
		RoleID:    role.ID,
	}
	if err := s.Add(&user); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

// Main check for api_key
func GetUserByKey(c *gin.Context){
	s := persistence.GetUserRepository()
	key := c.Params.ByName("api_key")
	// movie_id := c.Params.ByName("movie_id")
	if user, err := s.GetbyKey(key); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func FavMovie(c *gin.Context){
	user_repo := persistence.GetUserRepository()
	m_repo := persistence.GetMovieRepository()
	id := c.Params.ByName("id")
	user := c.MustGet("user").(*models.User)
	if movie, err := m_repo.Get(id); err != nil{
		http_err.NewError(c, http.StatusNotFound, errors.New("movie not found"))
		log.Println(err)
	}else{
			// Add movie to fav if user exists
			for _, item := range user.Favorites {
				if item.ID == movie.ID {
					fmt.Println("Movie already fav")
					http_err.NewError(c, http.StatusNotAcceptable, errors.New("Movie already favorited"))
					return
				}
			}
			user.Favorites = append(user.Favorites, movie)
			if err := user_repo.Update(user); err != nil {
				http_err.NewError(c, http.StatusNotFound, err)
				log.Println(err)
			} else {
				c.JSON(http.StatusOK, "User was upadated")
			}
	}
}

func ShowFavMovies(c *gin.Context){
	user := c.MustGet("user").(*models.User)
	c.JSON(http.StatusOK, user.Favorites)
}


func DeleteFavMovie(c *gin.Context){
	s := persistence.GetUserRepository()
	user := c.MustGet("user").(*models.User)
	id := c.Params.ByName("id")
	
}

func UpdateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Hash = crypto.GenerateHash([]byte(userInput.Password))

		if err := s.Update(user); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func DeleteUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			// TODO Succeed Message 
			c.JSON(http.StatusNoContent, "")
		}
	}
}