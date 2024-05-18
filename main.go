package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

var Movies = []Movie{
	{ID: "1", Title: "Toy Story", Genre: "Comedy"},
	{ID: "2", Title: "Grumpy Old Men Collection", Genre: "Romance"},
	{ID: "3", Title: "James Bond", Genre: "Action"},
	{ID: "4", Title: "Lapaata Ladies", Genre: "Romance"},
	{ID: "5", Title: "The Conjuring", Genre: "Horror"},
}

func getMovies(C *gin.Context) {
	C.IndentedJSON(http.StatusOK, Movies)

}
func main() {
	gin_route := gin.Default()
	gin_route.GET("/Movies", getMovies)
	gin_route.Run("localhost:8080")

}
