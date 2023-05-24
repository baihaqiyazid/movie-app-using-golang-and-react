package service

import (

	"server/cmd/helper"
	"server/cmd/models/entity"
	"server/cmd/repository"
	"server/cmd/web"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//================ IMPLEMENTATION =============================================

type MovieServiceImpl struct {
	DB *gorm.DB
	Repository repository.MovieRepository
}

//================ CONSTRUCTOR ================================================

func NewMovieService(db *gorm.DB, Repository repository.MovieRepository) *MovieServiceImpl {
	return &MovieServiceImpl{
		DB: db,
		Repository: Repository}
}

//================ METHOD =====================================================
func (service *MovieServiceImpl) GetAllMovies() (*[]entity.Movie, error) {
	movies, err := service.Repository.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (service *MovieServiceImpl) GetMovieById(id int) (*entity.Movie, error) {
	movies, err := service.Repository.GetMovieById(id)
	if err != nil {
		return nil, err
	}

	if movies == nil {
		return nil, nil
	}

	return movies, nil
}

func (service *MovieServiceImpl) GetMoviesByGenre(id int) (*[]entity.Genre, error) {
	genre, err := service.Repository.GetMoviesByGenre(id)
	if err != nil {
		return nil, err
	}
	return genre, nil
}

func (service *MovieServiceImpl) GetLastMovies() (*entity.Movie, error) {
	movie, err := service.Repository.GetLastMovies()
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (service *MovieServiceImpl) CreateMovie(request web.MoviePayloadResponse) error {
	var movieGenre entity.MovieGenre

	//create movie
	movie := movieGenre.Movie	
	movie.ID, _ = strconv.Atoi(request.ID)
	movie.Title = request.Title
	movie.Description = request.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", request.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(request.Runtime)
	movie.Rating, _ = strconv.Atoi(request.Rating)
	movie.MPAARating = request.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	err := service.Repository.CreateMovie(movie)
	if err != nil {
		return err
	}

	//get last movie
	movieLast, err := service.Repository.GetLastMovies()
	if err != nil {
		panic(err)
	}

	//create movie genre
	movieGenre.ID, _ = strconv.Atoi(request.ID)
	movieGenre.MovieID = movieLast.ID
	movieGenre.CreatedAt = time.Now()
	movieGenre.UpdatedAt = time.Now()
	for _, v := range request.GenreID {
		movieGenre.GenreID, _ = strconv.Atoi(v)
		err := service.Repository.CreateMovieGenres(movieGenre)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (service *MovieServiceImpl) UpdateMovie (request web.MoviePayloadResponse, ctx *gin.Context) error {
	
	id, _ := strconv.Atoi(request.ID)
	movie, err := service.Repository.GetMovieById(id)
	if err != nil {
		helper.NotFound(*ctx, err)
		return err
	}

	movie.ID, _ = strconv.Atoi(request.ID)
	movie.Title = request.Title
	movie.Description = request.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", request.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(request.Runtime)
	movie.Rating, _ = strconv.Atoi(request.Rating)
	movie.MPAARating = request.MPAARating
	movie.UpdatedAt = time.Now()

	err = service.Repository.UpdateMovie(movie, id)
	if err != nil {
		panic(err)
	}

	mg, _ := service.Repository.GetMovieGenresById(id)
	err = service.Repository.DeleteMovieGenres(mg.MovieID)
	if err != nil {
		panic(err)
	}
	mg.MovieID = movie.ID
	mg.UpdatedAt = time.Now()
	mg.CreatedAt = time.Now()
	mgLast, _ := service.Repository.GetLastMovieGenre()
	if (mgLast != nil){
		for index, v := range request.GenreID {
			mg.ID = mgLast.ID + index + 1
			mg.GenreID, _ = strconv.Atoi(v)
			err = service.Repository.CreateMovieGenres(*mg)
			if err != nil {
				panic(err)
			}
		}
	}else{
		for index, v := range request.GenreID {
			mg.GenreID, _ = strconv.Atoi(v)
			mg.ID = mg.ID + index
			err := service.Repository.CreateMovieGenres(*mg)
			if err != nil {
				panic(err)
			}
		}
	}
	
	return nil
}

func (service *MovieServiceImpl) DeleteMovie(request web.MoviePayloadResponse, ctx *gin.Context) error {

	id, _ := strconv.Atoi(request.ID)
	movie, err := service.Repository.GetMovieById(id)
	if err != nil {
		helper.NotFound(*ctx, err)
		return err
	}

	err = service.Repository.DeleteMovie(movie.ID)
	if err != nil {
		helper.NotFound(*ctx, err)
		return err
	}
	return nil
}

func (service *MovieServiceImpl) GetAllGenres() (*[]entity.Genre, error) {
	genres, err := service.Repository.GetAllGenres()
	if err != nil {
		return nil, err
	}
	return genres, nil
}