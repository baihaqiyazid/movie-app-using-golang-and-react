package app

import (
	"server/cmd/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(movieController controller.MovieController){
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	
	// router.GET("/status", app.StatusHandler)

	// MOVIES
	router.GET("/movies/:id", movieController.GetMovieById)
	router.GET("/movies", movieController.GetMovies)

	// GENRES
	router.GET("/genres", movieController.GetGenres)
	router.GET("/genres/:id", movieController.GetMoviesByGenre)

	// ADMIN
	router.POST("/admin/movies/create", movieController.CreateMovies)
	router.POST("/admin/movies/edit", movieController.UpdateMovie)
	router.DELETE("/admin/movies/delete", movieController.DeleteMovie)

	router.Run(":4001")
}