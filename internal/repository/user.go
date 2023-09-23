package repository

import (
	"auth_microservice/internal/domain"
	"context"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (r *Users) GetByCredentials(ctx context.Context, phoneNumber, password string) (domain.User, error) {
	var user domain.User
	selectStmt := "SELECT id FROM users WHERE phone_number=$1 AND password=$2"
	err := r.db.QueryRowContext(ctx, selectStmt, phoneNumber, password).
		Scan(&user.ID)

	return user, err
}
