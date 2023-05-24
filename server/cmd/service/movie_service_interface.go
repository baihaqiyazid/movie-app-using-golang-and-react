package service

import (
	"server/cmd/models/entity"
	"server/cmd/web"

	"github.com/gin-gonic/gin"
)

type MovieService interface {
	GetAllMovies() (*[]entity.Movie, error)
	GetMovieById(id int) (*entity.Movie, error)
	GetMoviesByGenre(id int) (*[]entity.Genre, error)
	GetLastMovies() (*entity.Movie, error)

	CreateMovie(request web.MoviePayloadResponse) error
	UpdateMovie(request web.MoviePayloadResponse, ctx *gin.Context) error
	DeleteMovie(request web.MoviePayloadResponse, ctx *gin.Context) error

	GetAllGenres() (*[]entity.Genre, error)
}