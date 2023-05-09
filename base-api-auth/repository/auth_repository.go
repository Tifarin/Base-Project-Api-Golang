package repository

import (
	"base-api-auth/data/model"
	"context"
)

type AuthRepository interface {
    FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.User, error)
}