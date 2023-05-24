package repository

import "server/cmd/models/entity"

type MovieRepository interface {
	GetAllMovies() (*[]entity.Movie, error)
	GetMovieById(id int) (*entity.Movie, error)
	GetMoviesByGenre(id int) (*[]entity.Genre, error)
	GetLastMovies() (*entity.Movie, error)

	CreateMovie(movie entity.Movie) error
	UpdateMovie(movie *entity.Movie, id int) error
	DeleteMovie(id int) error

	GetAllGenres() (*[]entity.Genre, error)

	GetMovieGenresById(id int) (*entity.MovieGenre, error)
	GetLastMovieGenre() (*entity.MovieGenre, error)

	CreateMovieGenres(movieGenre entity.MovieGenre) error
	UpdateMovieGenre(movieGenre *entity.MovieGenre, id int) error
	DeleteMovieGenres(id int) error
}
