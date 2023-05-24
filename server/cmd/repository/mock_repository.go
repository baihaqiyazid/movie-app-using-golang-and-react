package repository

import (
    "github.com/golang/mock/gomock"
)

func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
    return &MockMovieRepository{ctrl}
}

type MockMovieRepository struct {
    ctrl *gomock.Controller
}