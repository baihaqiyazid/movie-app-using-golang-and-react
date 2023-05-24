package test

import (
	"errors"
	"server/cmd/models/entity"

	"github.com/stretchr/testify/mock"
)

type mockMovieRepository struct {
	mock.Mock
}

// GetAllMovies
func (m *mockMovieRepository) GetAllMovies() (*[]entity.Movie, error) {
	args := m.Called()
	return args.Get(0).(*[]entity.Movie), args.Error(1)
}

// GetMoviesByID
func (m *mockMovieRepository) GetMovieById(id int) (*entity.Movie, error) {
	arguments := m.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Movie Not Found")
	}else {
		movie := arguments.Get(0).(*entity.Movie)
		return movie, nil
	}
} 

// GetMoviesByGenre
func (m *mockMovieRepository) GetMoviesByGenre(id int) (*[]entity.Genre, error) {
	arguments := m.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("Movie Not Found")
	}else {
		movie := arguments.Get(0).(*[]entity.Genre)
		return movie, nil
	}
} 

// GetAllGenres
func (m *mockMovieRepository) GetAllGenres() (*[]entity.Genre, error) {
	arguments := m.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("Genres Not Found")
	}else {
		genre := arguments.Get(0).(*[]entity.Genre)
		return genre, nil
	}
} 

// GetLastMovieGenre
func (m *mockMovieRepository) GetLastMovieGenre() (*entity.MovieGenre, error) {
	args := m.Called()
	return args.Get(0).(*entity.MovieGenre), args.Error(1)
} 

// GetMovieGenresById
func (m *mockMovieRepository) GetMovieGenresById(id int) (*entity.MovieGenre, error) {
	args := m.Called()
	return args.Get(0).(*entity.MovieGenre), args.Error(1)
} 

// CreateMovie
func (m *mockMovieRepository) CreateMovie(movie entity.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

// CreateMovieGenres
func (m *mockMovieRepository) CreateMovieGenres(movieGenre entity.MovieGenre) error {
	args := m.Called(movieGenre)
	return args.Error(0)
}

// GetLastMovies
func (m *mockMovieRepository) GetLastMovies() (*entity.Movie, error) {
	arguments := m.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("Movie Not Found")
	}else {
		movie := arguments.Get(0).(*entity.Movie)
		return movie, nil
	}
} 


func (m *mockMovieRepository) DeleteMovie(id int) error {
	return nil
} 

func (m *mockMovieRepository) DeleteMovieGenres(id int) error {
	return nil
} 


func (m *mockMovieRepository) UpdateMovie(movie *entity.Movie, id int) error {
	return nil
}

func (m *mockMovieRepository) UpdateMovieGenre(movie *entity.MovieGenre, id int) error {
	return nil
}