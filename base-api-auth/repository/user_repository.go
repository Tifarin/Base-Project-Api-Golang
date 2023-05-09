package repository

import (
	"base-api-auth/data/model"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, user model.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
}