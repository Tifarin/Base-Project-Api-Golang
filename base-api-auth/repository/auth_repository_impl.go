package repository

import (
	"base-api-auth/data/model"
	"context"
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.User, error) {
	query := `
		SELECT id, username, names, email, password_hash, id_roles, created_at, updated_at
		FROM users
		WHERE username = $1 OR email = $1
		LIMIT 1
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, usernameOrEmail).Scan(&user.ID, &user.Username, &user.Names, &user.Email, &user.PasswordHash, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}