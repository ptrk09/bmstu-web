package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int

	queryString := fmt.Sprintf(
		`INSERT INTO %s (name, login, password, role_id)  
		SELECT $1, $2, $3, id FROM roles WHERE name = $4`,
		usersTable,
	)
	row := r.db.QueryRow(queryString, user.Name, user.Login, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (model.User, error) {
	var user model.User

	queryString := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2;", usersTable)
	err := r.db.Get(&user, queryString, login, password)

	return user, err
}
