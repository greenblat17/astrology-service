package repo

import (
	"github.com/greenblat17/api/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type Star interface {
	FindAll() ([]domain.Star, error)
	FindByDate(date string) (domain.Star, error)
}

type Repository struct {
	Star
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Star: NewStarPostgres(db),
	}
}
