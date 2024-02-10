package service

import (
	"github.com/greenblat17/api/pkg/domain"
	"github.com/greenblat17/api/pkg/repo"
)

type StarService struct {
	repo repo.Star
}

func NewStarService(repo repo.Star) *StarService {
	return &StarService{repo: repo}
}

func (s *StarService) GetAllStars() ([]domain.Star, error) {
	return s.repo.FindAll()
}

func (s *StarService) GetStarByDate(date string) (domain.Star, error) {
	return s.repo.FindByDate(date)
}
