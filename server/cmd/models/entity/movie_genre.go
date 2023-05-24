package entity

import "time"

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	Movie     Movie     `gorm:"foreignKey:MovieID"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `gorm:"foreignKey:GenreID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}