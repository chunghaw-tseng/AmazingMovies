package main

import (
    "fmt"
    // "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "example.com/amazingmovies/src/pkg/db"
    "example.com/amazingmovies/src/api/controllers"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello my name!")
}


func main() {

	// http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
    //     log.Fatal(err)
    // }

    router := gin.Default()
    db.ConnectDatabase()

    // Process the templates at the start so that they don't have to be loaded
    // from the disk again. This makes serving HTML pages very fast.
    router.LoadHTMLGlob("static/*")

    // Define the route for the index page and display the index.html template
    // To start with, we'll use an inline route handler. Later on, we'll create
    // standalone functions that will be used as route handlers.
    router.GET("/", func(c *gin.Context) {
        // Call the HTML method of the Context to render a template
        c.HTML(
        // Set the HTTP status to 200 (OK)
        http.StatusOK,
        // Use the index.html template
        "index.html",
        // Pass the data that the page uses (in this case, 'title')
        gin.H{
            "title": "Home Page",
        },
        )
    })
    router.GET("/hello", controllers.TestFunction)
    router.GET("/movies", controllers.FindMovies)
  // Start serving the application
  router.Run()
}