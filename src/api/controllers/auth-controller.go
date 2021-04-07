package controllers

import(
	"log"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/crypto"
	"example.com/amazingmovies/pkg/http_errors"
)


type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


// Login godoc
// @Summary Retrieves JWT Token for specific user
// @Description Get JWT Token from creedentials
// @Produce string
// @Accept  json
// @Param username string
// @Param password string
// @Success 200 {string} JWT Token 
// @Router /am_api/login [post]
func Login(c *gin.Context) {
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		c.JSON(http.StatusOK, token)
	}
}

// KeyLogin godoc
// @Summary Retrieves API Key for specific user
// @Description Get API Key from creedentials
// @Produce string
// @Accept  json
// @Param username string
// @Param password string
// @Success 200 {string} APIKey
// @Router /am_api/loginkey [post]
func KeyLogin(c *gin.Context){
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		c.JSON(http.StatusOK, user.APIKey)
	}
	
}