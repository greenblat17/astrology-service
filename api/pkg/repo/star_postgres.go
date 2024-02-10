package repo

import (
	"github.com/greenblat17/api/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type StarPostgres struct {
	db *sqlx.DB
}

func NewStarPostgres(db *sqlx.DB) *StarPostgres {
	return &StarPostgres{db: db}
}

func (r *StarPostgres) FindAll() ([]domain.Star, error) {
	rows, err := r.db.Query("SELECT id, title, description, date, url FROM stars")
	if err != nil {
		return nil, err
	}

	var stars []domain.Star
	for rows.Next() {
		var star domain.Star
		if err := rows.Scan(&star.ID, &star.Name, &star.Description, &star.Date, &star.ImageURL); err != nil {
			return nil, err
		}
		stars = append(stars, star)
	}

	return stars, nil
}

func (r *StarPostgres) FindByDate(date string) (domain.Star, error) {
	row := r.db.QueryRow("SELECT id, title, description, date, url FROM stars WHERE date = $1", date)

	var star domain.Star
	if err := row.Scan(&star.ID, &star.Name, &star.Description, &star.ImageURL); err != nil {
		return star, err
	}

	return star, nil
}
