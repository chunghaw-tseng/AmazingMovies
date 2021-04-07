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


type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct{
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
}


// GetUserById godoc
// @Summary get specific user from db
// @Description get user with specific id
// @Produce json
// @Param id integer "User.ID"
// @Success 200 {object} model.User
// @Router /am_api/users/{id} [get]
// @Security Admin Authorization Token 
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


// GetUsers godoc
// @Summary Fetch all the users
// @Description get all the users from db
// @Produce json
// @Success 200 {object} model.User
// @Router /am_api/users [get]
// @Security Admin Authorization Token 
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

// CreateUser godoc
// @Summary Creates new user
// @Description adds new user to db
// @Accepts json
// @Param username string
// @Param lastname string
// @Param username string
// @Param password string
// @Produce json
// @Success 200 {object} model.User
// @Router /am_api/users [put]
func CreateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	r := persistence.GetRolesRepository()	
	var userInput CreateUserInput
	_ = c.BindJSON(&userInput)
	apikey := uuid.New().String()
	role, _ := r.Get("user")
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



// FavMovie godoc
// @Summary Adds movie to user favorites
// @Description  Adds new movie to the user list of favorites
// @Param id integer "Movie.ID"
// @Produce json
// @Success 200 {object} string
// @Failure 406 {object} Error
// @Failure 404 {object} Error
// @Router /am_api/favorite/{id} [post]
// @Authentication User Authentication Token or API Key
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

// ShowFavMovies godoc
// @Summary show the list of favorite movies
// @Description  gets all the movies that are favorited
// @Produce json 
// @Success 200 {object} []model.Movies
// @Router /am_api/favorite [get]
// @Authentication User Authentication Token or API Key
func ShowFavMovies(c *gin.Context){
	user := c.MustGet("user").(*models.User)
	c.JSON(http.StatusOK, user.Favorites)
}


// DeleteFavMovie godoc
// @Summary Delete movie from favorites
// @Description  removes the movie association with from the favorites list
// @Param id integer "Movie.ID"
// @Produce json
// @Success 200 
// @Failure 404 {object} Error
// @Router /am_api/favorite/{id} [delete]
// @Authentication User Authentication Token or API Key
func DeleteFavMovie(c *gin.Context){
	s := persistence.GetUserRepository()
	m_repo := persistence.GetMovieRepository()
	user := c.MustGet("user").(*models.User)
	id := c.Params.ByName("id")
	if movie, err := m_repo.Get(id); err != nil{
		http_err.NewError(c, http.StatusNotFound, errors.New("movie not found"))
		log.Println(err)
	}else{
		if err := s.DeleteAssociation(user, "Favorites", movie); err != nil{
			log.Println(err)
		}
	}

}


// UpdateUser godoc
// @Summary Updates the user info
// @Description  updates the user information
// @Produce json
// @Param lastname string
// @Param firstname string
// @Param password string
// @Success 200 {object} models.User
// @Failure 404 {object} Error
// @Router /am_api/users [put]
// @Authentication User Authentication Token or API Key
func UpdateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	user := c.MustGet("user").(*models.User)
	var userInput UpdateUserInput
	_ = c.BindJSON(&userInput)
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


// DeleteUser godoc
// @Summary Delete specific user
// @Description  removes the user with specific id
// @Param id integer "User.ID"
// @Produce json
// @Success 200 
// @Failure 404 {object} Error
// @Router /am_api/users/{id} [delete]
// @Authentication Admin Authentication Token
func DeleteUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	if user, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}