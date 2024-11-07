package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	IMDb  float64 `json:"IMDb"`
	Actor Actor   `json:"actor"`
}

type Actor struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies = []Movies{}

func getMovies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, movies)
}

func getMovieById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	fmt.Println(id)
	if err != nil {
		return
	}
	for _, movie := range movies {
		if movie.Id == id {
			ctx.IndentedJSON(http.StatusOK, movie)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found!"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	movies = append(movies, Movies{Id: 1, Title: "Iron Man", IMDb: 9.0, Actor: Actor{Firstname: "Robert", LastName: "Downey"}})
	movies = append(movies, Movies{Id: 2, Title: "Thor", IMDb: 8.6, Actor: Actor{Firstname: "Chris", LastName: "Hemsworth"}})
	movies = append(movies, Movies{Id: 3, Title: "Batman", IMDb: 9.2, Actor: Actor{Firstname: "Christian", LastName: "Bale"}})

	r.GET("/getmovies", getMovies)
	r.GET("/getmovie/:id", getMovieById)

	r.Run()
}
