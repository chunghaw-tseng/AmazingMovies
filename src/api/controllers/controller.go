package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


// TestFunction godoc
// @Summary Hello world test function
// @Description returns hello world
// @Produce json
// @Success 200 {object} json
// @Router /am_api/hello [get]
func TestFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
  }
  

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",
        gin.H{
            "title": "Home Page",
        },
        )
  }