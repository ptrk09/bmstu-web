package repository

import (
	"fmt"
	"web/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db}
}

func (r *UserPostgres) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf(
		`SELECT 
		u.id AS id, u.name AS name, 
		u.login AS login, u.password AS password, 
		r.name AS role
	FROM %s u JOIN %s r ON u.role_id = r.id`,
		usersTable, rolesTable,
	)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgres) GetUserById(userId int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf(
		`SELECT 
			u.id AS id, u.name AS name, 
			u.login AS login, u.password AS password, 
			r.name AS role
		FROM %s u JOIN %s r ON u.role_id = r.id
		WHERE u.id = 1;`,
		usersTable, rolesTable,
	)
	err := r.db.Get(&user, query)

	return user, err
}

func (r *UserPostgres) UpdateUser(userId int, input model.UpdateUserInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	updateUserQuery := fmt.Sprintf(
		`UPDATE %s
		SET name = $1, password = $2
		WHERE id = $3
		RETURNING id;`,
		usersTable,
	)

	row := tx.QueryRow(updateUserQuery, input.Name, input.Password, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *UserPostgres) DeleteUser(userId int) (int, error) {
	return 0, nil
}
