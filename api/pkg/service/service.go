package service

import (
	"github.com/greenblat17/api/pkg/domain"
	"github.com/greenblat17/api/pkg/repo"
)

type Star interface {
	GetAllStars() ([]domain.Star, error)
	GetStarByDate(date string) (domain.Star, error)
}

type Service struct {
	Star
}

func NewService(r *repo.Repository) *Service {
	return &Service{
		Star: NewStarService(r.Star),
	}
}
