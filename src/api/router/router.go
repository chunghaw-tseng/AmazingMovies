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
	// TODO Get API -> Change to accounts
	app.POST("/am_api/users", controllers.CreateUser)
	app.DELETE("/am_api/users/:id", controllers.DeleteUser)


	// Only Admin
	app.GET("am_api/users", controllers.GetUsers)
	app.GET("am_api/users/id/:id", controllers.GetUserById)
	app.GET("am_api/users/key/:api_key", controllers.GetUserByKey)
	app.PUT("/am_api/users/:id", controllers.UpdateUser)

	// ========== Usage
	// Movies
	app.GET("/am_api/movies", controllers.GetMovies)
    app.GET("/am_api/movies/:id", controllers.GetMoviesById)
	app.POST("/am_api/movies", controllers.CreateMovie)
	// app.PUT("/api/movies/:id", controllers.)
	// app.DELETE("/api/movies/:id", controllers.DeleteTask)

	// Genre
	app.GET("/am_api/genres", controllers.GetGenres)
	app.POST("/am_api/genres", controllers.CreateGenre)

	// ================== Login Routes
	// app.POST("/api/login", controllers.Login)
	// app.POST("/api/login/add", controllers.CreateUser)
	// app.GET("/api/users", controllers.GetUsers)
	// app.GET("/api/users/:id", controllers.GetUserById)
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