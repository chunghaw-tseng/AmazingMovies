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

	// TODO Add authentication
	// app.Use(middlewares.CORS())
	app.NoRoute(middleware.NoRouteHandler())

	// ========= Static Routes
	app.LoadHTMLGlob("static/*")

	// TODO
	// ========== Docs Routes
	// app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// ========= Test Routes
    app.GET("/", controllers.ShowIndex)
    app.GET("/hello", controllers.TestFunction)
    

	// ========= Login Routes


	// ========== User Routes
	// TODO Get API
	// app.GET("/favorites", controllers)
	// app.POST("/favorits", )


	// ========== Usage
	// Movies
	app.GET("/api/movies", controllers.GetMovies)
    app.GET("/api/movies/:id", controllers.GetMoviesById)
	app.POST("/api/movies", controllers.CreateMovie)
	// app.PUT("/api/movies/:id", controllers.)
	// app.DELETE("/api/movies/:id", controllers.DeleteTask)

	// Genre
	app.GET("/api/genres", controllers.GetGenres)
	app.POST("/api/genres", controllers.CreateGenre)

	// ================== Login Routes
	// app.POST("/api/login", controllers.Login)
	// app.POST("/api/login/add", controllers.CreateUser)
	// app.GET("/api/users", controllers.GetUsers)
	// app.GET("/api/users/:id", controllers.GetUserById)
	// app.POST("/api/users", controllers.CreateUser)
	// app.PUT("/api/users/:id", controllers.UpdateUser)
	// app.DELETE("/api/users/:id", controllers.DeleteUser)
	// ================== Tasks Routes
	// app.GET("/api/tasks/:id", controllers.GetTaskById)
	// app.GET("/api/tasks", controllers.GetTasks)
	// app.POST("/api/tasks", controllers.CreateTask)
	// app.PUT("/api/tasks/:id", controllers.UpdateTask)
	// app.DELETE("/api/tasks/:id", controllers.DeleteTask)

	return app
}