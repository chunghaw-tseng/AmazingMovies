package middleware

import (
	"example.com/amazingmovies/pkg/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)


// Checks and depends on the API will check any of the bottom
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		// TODO Check if authorization is a API Key or a token
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}


func TokenRequired(){

}

// TODO Check the API
func APIKeyRequired(){

}
