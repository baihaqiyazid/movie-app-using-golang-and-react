package controller

import (
	"log"
	"strconv"

	"server/cmd/helper"
	"server/cmd/service"
	"server/cmd/web"

	"github.com/gin-gonic/gin"
)

type MovieController interface{
	GetMovies(ctx *gin.Context)
	GetMovieById(ctx *gin.Context)
	GetMoviesByGenre(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	GetGenres(ctx *gin.Context)
	CreateMovies(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
}

type MovieControllerImpl struct{
	MovieService service.MovieService
}

func NewMovieController(service service.MovieService) *MovieControllerImpl{
	return &MovieControllerImpl{MovieService: service}
}

func (controller *MovieControllerImpl) GetMovies(ctx *gin.Context) {

	movie, err := controller.MovieService.GetAllMovies()
	helper.PanicError(err)

	helper.ResponseSuccess(ctx, movie)
}

func (controller *MovieControllerImpl) GetMovieById(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		helper.BadRequest(*ctx, err)
		return
	}

	movie, err := controller.MovieService.GetMovieById(id)
	if err != nil {
		helper.NotFound(*ctx, err)
		return
	}
	helper.ResponseSuccess(ctx, movie)
}

func (controller *MovieControllerImpl) GetMoviesByGenre(ctx *gin.Context) {

	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		helper.BadRequest(*ctx, err)
		return
	}

	movie, err := controller.MovieService.GetMoviesByGenre(id)
	if err != nil {
		helper.NotFound(*ctx, err)
		return
	}

	helper.ResponseSuccess(ctx, movie)
}

func (controller *MovieControllerImpl) DeleteMovie(ctx *gin.Context) {

	var request web.MoviePayloadResponse
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	err := controller.MovieService.DeleteMovie(request, ctx)
	if err != nil {
		helper.NotFound(*ctx, err)
		return
	}
	helper.ResponseSuccess(ctx, nil)
}

func (controller *MovieControllerImpl) GetGenres(ctx *gin.Context) {

	genres, err := controller.MovieService.GetAllGenres()
	helper.PanicError(err)
	helper.ResponseSuccess(ctx, genres)
}

func (controller *MovieControllerImpl) CreateMovies(ctx *gin.Context) {

	var request web.MoviePayloadResponse
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}
	
	err := controller.MovieService.CreateMovie(request)
	helper.PanicError(err)

	helper.ResponseSuccess(ctx, nil)

	return
}

func (controller *MovieControllerImpl) UpdateMovie(ctx *gin.Context) {
	var request web.MoviePayloadResponse
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	err := controller.MovieService.UpdateMovie(request, ctx)
	helper.PanicError(err)
	helper.ResponseSuccess(ctx, nil)

	return
}

