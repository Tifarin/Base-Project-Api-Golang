package service

import (
	"base-api-auth/data/model"
	"context"
)

type AuthService interface {
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.User, error)
}