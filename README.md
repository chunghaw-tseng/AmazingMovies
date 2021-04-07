# AmazingMovies

API Server written in GoLang for amazing movies
There are some data that will be initialized for easier testing.
Admin user is created by default {username:admin, passsword:admin}

Decided to create this backend to learn more and expand my knowledge on golang.
I don't have much experience with this language yet.

### Usage

- To start run after initializing all the go modules

```
    go run main.go
```

- You can access the API Key or the JWT Token with the login and loginkey apis
- To use JWT please add Bearer in front in the Authentication Header
- Some documentation is written for some APIs
- There are 3 kinds of APIs

1. Need User Bearer Token/API Key

- Update User, Add Favorite, Delete Favorite and See Favorites

2. Need Bearer Token from an Admin Role user

- Get Users, Delete Users, Delete Movies

3. Open to anyone (No need for API Key/ Bearer Token or Admin Role)

- The rest

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

- Deletion of some data such as genre and people
- Use ginswagger for Documentation
