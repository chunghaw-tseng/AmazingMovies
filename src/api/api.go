package api

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"example.com/amazingmovies/src/pkg/db"
	"example.com/amazingmovies/src/pkg/config"
	"example.com/amazingmovies/src/api/router"

)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.StartDatabase()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}