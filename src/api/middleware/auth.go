package middleware

import (
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
		s := persistence.GetUserRepository()
		if authorizationHeader != "" {
			// Token need to return user info
			if strings.Contains(authorizationHeader, "Bearer"){
				auth_slice := strings.Fields(authorizationHeader)
				if valid, username := crypto.ValidateToken(auth_slice[1]); !valid{
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : token not accepted"})
					return
				} else {
					if user , err := s.GetByUsername(username); err != nil {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : token not accepted"})
					} else {
						c.Set("user", user)
						c.Next()
					}
				}
			}else{
					// API Key
					if user, err := s.GetbyKey(authorizationHeader); err != nil {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : api_key not valid"})
					} else {
						c.Set("user", user)
						c.Next()
					}
				}
		}else{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header found"})
		}
	}
}

// Checks if user is admin for deleting
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the token is a Bearer token
		authorizationHeader := strings.TrimSpace(c.GetHeader("authorization"))
		s := persistence.GetUserRepository()
		if authorizationHeader != "" {
			// Token need to return user info
			if strings.Contains(authorizationHeader, "Bearer"){
				auth_slice := strings.Fields(authorizationHeader)
				if valid, username := crypto.ValidateToken(auth_slice[1]); !valid{
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : token not accepted"})
					return
				} else {
					if user , err := s.GetByUsername(username); err != nil {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : token not accepted"})
					} else {
						if user.RoleID == 1{
							c.Set("user", user)
							c.Next()
							return
						}else{
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized : Admin Role Required"})
						}
					}
				}
			}else{
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Bearer Token Found"})
			}
		}else{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header found"})
		}
	}
}