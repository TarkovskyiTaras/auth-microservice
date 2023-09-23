package repository

import (
	"auth_microservice/internal/domain"
	"context"
	"database/sql"
)

type Tokens struct {
	db *sql.DB
}

func NewTokens(db *sql.DB) *Tokens {
	return &Tokens{db}
}

func (r *Tokens) Create(ctx context.Context, token domain.Token) error {
	createStmt := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, createStmt, token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (r *Tokens) Get(ctx context.Context, token string) (domain.Token, error) {
	var t domain.Token
	selectStmt := "SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1"
	err := r.db.QueryRowContext(ctx, selectStmt, token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)
	if err != nil {
		return t, err
	}

	deleteStmt := "DELETE FROM refresh_tokens WHERE user_id=$1"
	_, err = r.db.ExecContext(ctx, deleteStmt, t.UserID)

	return t, err
}
