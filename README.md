# AmazingMovies

API Server written in GoLang for amazing movies

## Libraries used

gin-swagger (Documentation)
jwt-go (WebToken for Auths)
gin (HTTP API) / Other option is Revel (But this seems to be not that popular)
gorm (Database SQL)
viper (ENV and Command line stuff)

## Purpose

Build an API, that supports:
A. getting a list of movies filterable by a query and B. details for a specific movie.
C. favorite a specific movie.
D. get a list of favorited movies.

Your API should expose the following endpoints:
GET /favorites - return that the user has previously favoritted.
POST /favorite/:id - add a favorite movie
GET /movies?search={search} - return popular movies or what the user searched for GET /movies/:id - return that specific movie in detail
This functionality should be achieved by supporting the following endpoints:
GET /movies - This should take apiKey and searchQuery as parameters and return a well structured JSON collection of matching movies.
GET /movies/:id - This should take apiKey as a parameter and return full details of the requested movie in well structured JSON.
The rest is handled locally by your API.

### TODO

- Login to app to save favourite
- DB with some movies and data about 10 hard coded stuff
- Accept API Key with Authentication header
- Add authentication middleware and use a bearer token every time to authenticate this
- Cannot create same username

## Done

- Started a very easy implementation of the JWTs to see and learn how Golang achieves this. This method is still not complete and it needs more security but due to time restrictions, I will not go very deep into this area for now.
