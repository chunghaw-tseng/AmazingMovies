package router

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"example.com/amazingmovies/src/api/controllers"
	"example.com/amazingmovies/src/api/middleware"
	"io"
	"os"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// For panics and will return 500 if there is one
	app.Use(gin.Recovery())

	app.NoRoute(middleware.NoRouteHandler())

	// ========= Static Routes
	app.LoadHTMLGlob("static/*")


	// ========= Test Routes
    app.GET("/", controllers.ShowIndex)
    app.GET("/hello", controllers.TestFunction)
	// ========= Login Routes
	app.POST("/login", controllers.Login)
	app.POST("/loginkey", controllers.KeyLogin)

	am_api := app.Group("/am_api")
	{	
	
		// ========== User Routes
		am_api.POST("/users", controllers.CreateUser)
		am_api.PUT("/users", middleware.AuthRequired(),controllers.UpdateUser)
		// Favorites
		am_api.POST("/favorite/:id", middleware.AuthRequired() ,controllers.FavMovie)
		am_api.GET("/favorite", middleware.AuthRequired() ,controllers.ShowFavMovies)
		am_api.DELETE("/favorite/:id", middleware.AuthRequired() ,controllers.DeleteFavMovie)

		// Only Admin
		am_api.DELETE("/users/:id", middleware.AdminRequired(), controllers.DeleteUser)
		am_api.GET("/users/:id", middleware.AdminRequired(), controllers.GetUserById)
		am_api.GET("/users", middleware.AdminRequired(), controllers.GetUsers)

		// ========== Movies Usage
		am_api.GET("/movies", controllers.GetMovies)
		am_api.GET("/movies/:id", controllers.GetMoviesById)


		// =========== Adding Data
		am_api.POST("/movies", controllers.CreateMovie)
		am_api.PUT("/movies/:id", controllers.UpdateMovie)
		am_api.DELETE("/movies/:id", middleware.AdminRequired(), controllers.DeleteMovie)

		// Genre
		am_api.GET("/genres", controllers.GetGenres)

		// People
		am_api.GET("/people", controllers.GetPeople)
		am_api.PUT("/people/:id", controllers.UpdatePeople)

	}
	

	return app
}