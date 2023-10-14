package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/inzarubin80/booking-services"
)

type AuthMssql struct {
	db *sqlx.DB
}

func NewAuthMssql(db *sqlx.DB) *AuthMssql {
	return &AuthMssql{db: db}
}

func (r *AuthMssql) CreateUser(user booking.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthMssql) GetUser(username, password string) (booking.User, error) {
	var user booking.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}