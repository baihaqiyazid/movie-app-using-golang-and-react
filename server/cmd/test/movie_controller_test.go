package test

import (
	"fmt"
	"log"
	"server/cmd/app"
	"server/cmd/controller"
	"server/cmd/models/entity"
	"server/cmd/repository"
	"server/cmd/service"
	"server/cmd/web"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/testgoreactmovies?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func setupRouter(db *gorm.DB){
	
	db, _ = setupDB()

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(db, movieRepository)
	movieController := controller.NewMovieController(movieService)

	app.Route(movieController)
}

var mockRepo = &mockMovieRepository{}
var mockService = &service.MovieServiceImpl{Repository: mockRepo}

func truncateDB(db *gorm.DB)  {
	db.Exec("TRUNCATE movies")
	db.Exec("TRUNCATE movie_genres")
	db.Exec("TRUNCATE genres")
}

// GetAllMovies
func TestGetAllMoviesFound(t *testing.T) {
	expectedMovies := []entity.Movie{
		{ID: 1, Title: "Movie 1", Description: "Description 1"},
		{ID: 2, Title: "Movie 2", Description: "Description 2"},
	}

	// set up the expectations for the mock repository
	mockRepo.On("GetAllMovies").Return(&expectedMovies, nil)

	// call the service method
	movies, err := mockService.GetAllMovies()

	// assert the result
	assert.Nil(t, err)
	assert.NotNil(t, movies)
	for i := range expectedMovies {
		fmt.Println("Expected ID: " , expectedMovies[i].ID,  "Actual ID: ", (*movies)[i].ID)
		fmt.Println("Expected ID: " , expectedMovies[i].Description, "Actual ID: ", (*movies)[i].Description)

        assert.Equal(t, expectedMovies[i].ID, (*movies)[i].ID)
    }
	mockRepo.AssertExpectations(t)
}

// GetMovieByID
func TestGetMovieByIdFound(t *testing.T) {
	expectedMovies := &entity.Movie{
		ID: 2, 
		Title: "Movie 2", 
		Description: "Description 2",
	}

	// set up the expectations for the mock repository
	mockRepo.On("GetMovieById",2).Return(expectedMovies)
	
	// call the service method
	movies, err := mockService.GetMovieById(2)

	// assert the result
	assert.Nil(t, err)
	assert.NotNil(t, movies)
	fmt.Println("Expected ID: " , expectedMovies.ID,  "Actual ID: ", movies.ID)
	fmt.Println("Expected Title: " , expectedMovies.Title, "Actual Title: ", movies.Title)
	fmt.Println("Expected Description: " , expectedMovies.Description, "Actual Description: ", movies.Description)

	assert.Equal(t, expectedMovies.ID, movies.ID)

	mockRepo.AssertExpectations(t)
}

// GetMovieByID (NOT FOUND)
func TestGetMovieByIdNotFound(t *testing.T) {
	// set up the expectations for the mock repository
	mockRepo.On("GetMovieById",0).Return(nil)
	
	// call the service method
	movies, err := mockService.GetMovieById(0)

	// assert the result
	assert.Nil(t, movies)
	assert.NotNil(t, err)
	fmt.Println(movies, err)
}

// GetMoviesByGenre
func TestGetMoviesByGenre(t *testing.T) {
	expectedMovies := []entity.Genre{
		{ID: 1, Movies: []entity.Movie{{ID: 1}}},
		{ID: 1,  Movies: []entity.Movie{{ID: 2}}},
	}

	// set up the expectations for the mock repository
	mockRepo.On("GetMoviesByGenre", 1).Return(&expectedMovies, nil)

	// call the service method
	movies, err := mockService.GetMoviesByGenre(1)

	// assert the result
	assert.Nil(t, err)
	assert.NotNil(t, movies)
	for i := range expectedMovies {
		fmt.Println("Expected ID: " , expectedMovies[i].ID,  "Actual ID: ", (*movies)[i].ID)
		for j := range expectedMovies[i].Movies {
			fmt.Println("Expected Movies ID: " , expectedMovies[i].Movies[j].ID, "Actual Movies ID: ", (*movies)[i].Movies[j].ID)
			assert.Equal(t, expectedMovies[i].Movies[j].ID, (*movies)[i].Movies[j].ID)
		}
		assert.Equal(t, expectedMovies[i].ID, (*movies)[i].ID)
		assert.Equal(t, expectedMovies[i].GenreName, (*movies)[i].GenreName)
    }
	mockRepo.AssertExpectations(t)
}

// GetMoviesByGenre (NOT FOUND)
func TestGetMoviesByGenreNotFound(t *testing.T) {

	// set up the expectations for the mock repository
	mockRepo.On("GetMoviesByGenre", 1).Return(nil)

	// call the service method
	movies, err := mockService.GetMoviesByGenre(1)

	// assert the result
	assert.Nil(t, movies)
	assert.NotNil(t, err)
	fmt.Println(movies, err)

	mockRepo.AssertExpectations(t)
}

// GetAllGenres
func TestGetAllGenresFound(t *testing.T) {
	expectedGenres := []entity.Genre{
		{ID: 1, GenreName: "Animation"},
		{ID: 2, GenreName: "Action"},
	}

	// set up the expectations for the mock repository
	mockRepo.On("GetAllGenres").Return(&expectedGenres, nil)

	// call the service method
	genres, err := mockService.GetAllGenres()

	// assert the result
	assert.Nil(t, err)
	assert.NotNil(t, genres)
	for i := range expectedGenres {
		fmt.Println("Expected ID: " , expectedGenres[i].ID,  "Actual ID: ", (*genres)[i].ID)
		fmt.Println("Expected Genre Name: ", expectedGenres[i].GenreName, "Actual Genre Name: ", (*genres)[i].GenreName)

        assert.Equal(t, expectedGenres[i].ID, (*genres)[i].ID)
		assert.Equal(t, expectedGenres[i].GenreName, (*genres)[i].GenreName)
    }
	mockRepo.AssertExpectations(t)
}

// GetAllGenres (Not Found)
func TestGetAllGenresNotFound(t *testing.T) {
	// set up the expectations for the mock repository
	mockRepo.On("GetAllGenres").Return(nil)

	// call the service method
	genres, err := mockService.GetAllGenres()

	// assert the result
	assert.Nil(t, genres)
	assert.NotNil(t, err)
	fmt.Println(genres, err)
	mockRepo.AssertExpectations(t)
}

// GetLastMovie
func TestGetLastMovie(t *testing.T) {
	expectedMovies := &entity.Movie{
		ID: 2, 
		Title: "Movie 2", 
		Description: "Description 2",
	}

	// set up the expectations for the mock repository
	mockRepo.On("GetLastMovies").Return(expectedMovies)
	
	// call the service method
	movies, err := mockService.GetLastMovies()

	// assert the result
	assert.Nil(t, err)
	assert.NotNil(t, movies)
	fmt.Println("Expected ID: " , expectedMovies.ID,  "Actual ID: ", movies.ID)
	fmt.Println("Expected Title: " , expectedMovies.Title, "Actual Title: ", movies.Title)
	fmt.Println("Expected Description: " , expectedMovies.Description, "Actual Description: ", movies.Description)

	assert.Equal(t, expectedMovies.ID, movies.ID)
	assert.Equal(t, expectedMovies.Title, movies.Title)
	assert.Equal(t, expectedMovies.Description, movies.Description)

	mockRepo.AssertExpectations(t)
}

// GetLastMovie (NOT FOUND)
func TestGetLastMovieNotFound(t *testing.T) {

	// set up the expectations for the mock repository
	mockRepo.On("GetLastMovies").Return(nil)
	
	// call the service method
	movies, err := mockService.GetLastMovies()

	// assert the result
	assert.Nil(t, movies)
	assert.NotNil(t, err)
	fmt.Println(movies, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateMovie(t *testing.T) {
	db, _ := setupDB()
	truncateDB(db)
	setupRouter(db)

	mockRepo.On("CreateMovie", mock.AnythingOfType("*entity.Movie")).Return(nil)
	mockRepo.On("GetLastMovies").Return(&entity.Movie{ID: 1}, nil)
	mockRepo.On("CreateMovieGenres", mock.AnythingOfType("*entity.MovieGenre")).Return(nil)

	// set up the service
	mockService := &service.MovieServiceImpl{Repository: mockRepo}

	// set up the test request
	request := web.MoviePayloadResponse{
		ID:           "1",
		Title:        "Test Movie",
		Description:  "Test Description",
		ReleaseDate:  "2020-01-01",
		Runtime:      "120",
		Rating:       "8",
		MPAARating:   "R",
		GenreID:      []string{"1", "2"},
	}

	// call the service method
	err := mockService.CreateMovie(request)

	// assert the result
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}