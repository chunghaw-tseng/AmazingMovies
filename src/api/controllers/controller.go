package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


// TestFunction godoc
// @Summary Retrieves task based on given ID
// @Description get Task by ID
// @Produce json
// @Param id path integer true "Task ID"
// @Success 200 {object} tasks.Task
// @Router /api/tasks/{id} [get]
// @Security Authorization Token
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