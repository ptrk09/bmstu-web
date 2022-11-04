package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func NewTeamPostgres(db *sqlx.DB) *TeamPostgres {
	return &TeamPostgres{db}
}

func (r *TeamPostgres) Create(userId int, team model.Team) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var teamId int
	query := fmt.Sprintf("INSERT INTO %s (name, owner_id) values ($1, $2) RETURNING id", teamsTable)

	row := tx.QueryRow(query, team.Name, userId)
	err = row.Scan(&teamId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return teamId, tx.Commit()
}

func (r *TeamPostgres) GetAll(userId int) ([]model.Team, error) {
	var teams []model.Team

	query := fmt.Sprintf("SELECT * FROM %s", teamsTable)
	err := r.db.Select(&teams, query)

	return teams, err
}
