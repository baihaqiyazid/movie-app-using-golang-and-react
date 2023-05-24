package repository

import (
	"server/cmd/models/entity"
	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepositoryImpl{
	return &MovieRepositoryImpl{DB: db}
}

//------------------------------------------ MOVIES
func (MovieRepositoryImpl *MovieRepositoryImpl) GetAllMovies() (*[]entity.Movie, error) {
	var movies []entity.Movie
	err := MovieRepositoryImpl.DB.Preload("Genres").Order("title").Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return &movies, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) GetMovieById(id int) (*entity.Movie, error) {
	var movie entity.Movie
	err := MovieRepositoryImpl.DB.Preload("Genres").First(&movie, id).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) GetMoviesByGenre(id int) (*[]entity.Genre, error) {
	var genre []entity.Genre
	err := MovieRepositoryImpl.DB.Preload("Movies").First(&genre, id).Error
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) GetLastMovies() (*entity.Movie, error) {
	var movies entity.Movie
	err := MovieRepositoryImpl.DB.Last(&movies).Error
	if err != nil {
		return nil, err
	}

	//log.Println(movies)

	return &movies, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) CreateMovie(movie entity.Movie) error {
	err := MovieRepositoryImpl.DB.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) UpdateMovie(movie *entity.Movie, id int) error {
	err := MovieRepositoryImpl.DB.Where("id = ?", id).Save(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) DeleteMovie(id int) error {
	var movie entity.Movie
	err := MovieRepositoryImpl.DeleteMovieGenres(id)
	if err != nil {
		return err
	}

	err = MovieRepositoryImpl.DB.Delete(&movie, id).Error
	if err != nil {
		return  err
	}

	return nil
}

//------------------------------------------- GENRES
func (MovieRepositoryImpl *MovieRepositoryImpl) GetAllGenres() (*[]entity.Genre, error) {
	var genres []entity.Genre

	err := MovieRepositoryImpl.DB.Order("genre_name").Find(&genres).Error
	if err != nil {
		return nil, err
	}

	return &genres, nil
}

//------------------------------------------- MOVIE GENRES
func (MovieRepositoryImpl *MovieRepositoryImpl) GetMovieGenresById(id int) (*entity.MovieGenre, error) {
	var movie entity.MovieGenre
	err := MovieRepositoryImpl.DB.Where("movie_id = ? ", id).First(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) GetLastMovieGenre() (*entity.MovieGenre, error) {
	var movie entity.MovieGenre
	err := MovieRepositoryImpl.DB.Last(&movie).Error
	if err != nil {
		return nil, err
	}

	//log.Println(movie)

	return &movie, nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) CreateMovieGenres(movieGenre entity.MovieGenre) error {
	err := MovieRepositoryImpl.DB.Create(&movieGenre).Error
	if err != nil {
		return err
	}
	return nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) UpdateMovieGenre(movieGenre *entity.MovieGenre, id int) error {
	err := MovieRepositoryImpl.DB.Where("movie_id = ?", id).Save(&movieGenre).Error
	if err != nil {
		return err
	}
	return nil
}

func (MovieRepositoryImpl *MovieRepositoryImpl) DeleteMovieGenres(id int) error {
	var movieGenre entity.MovieGenre
	err := MovieRepositoryImpl.DB.Where("movie_id = ?", id).Delete(&movieGenre).Error
	if err != nil {
		return err
	}

	return nil
}

