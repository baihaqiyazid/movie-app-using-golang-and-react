package entity

import "time"

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"name"`
	Movies    []Movie   `json:"movies" gorm:"many2many:movie_genres"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}