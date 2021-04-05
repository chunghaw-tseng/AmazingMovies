package middleware

import (
	"fmt"
	"example.com/amazingmovies/src/pkg/persistence"
	"example.com/amazingmovies/pkg/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// Checks and depends on the API will check any of the bottom
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the token is a Bearer token
		authorizationHeader := strings.TrimSpace(c.GetHeader("authorization"))
		
		
		if authorizationHeader != "" {
			// Token need to return user info
			if strings.Contains(authorizationHeader, "Bearer"){
				auth_slice := strings.Fields(authorizationHeader)
				if !crypto.ValidateToken(auth_slice[1]) {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : token not accepted"})
					return
				} else {
					c.Next()
				}
			}else{
					// API Key
					fmt.Println("API Key ", authorizationHeader)
					s := persistence.GetUserRepository()
					if user, err := s.GetbyKey(authorizationHeader); err != nil {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : api_key not valid"})
					} else {
						c.Set("UserID", user.ID)
						c.Next()
					}
				}
		}else{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "need authorization header"})
		}
	}
}