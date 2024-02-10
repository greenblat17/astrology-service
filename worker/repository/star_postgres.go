package repository

import (
	"log"

	"github.com/greenblat17/worker/domain"
	"github.com/jmoiron/sqlx"
)

type StarPostgres struct {
	db *sqlx.DB
}

func NewStarPostgres(db *sqlx.DB) *StarPostgres {
	return &StarPostgres{db: db}
}

func (r *StarPostgres) Save(star domain.Star) error {
	log.Print("Insert Query", star.Name)
	query := "INSERT INTO stars (title, description, date, url) VALUES ($1, $2, $3, $4)"
	if _, err := r.db.Exec(query, star.Name, star.Description, star.Date, star.ImageURL); err != nil {
		return err
	}
	return nil
}
