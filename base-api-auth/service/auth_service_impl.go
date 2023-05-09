package service

import (
	"base-api-auth/repository"
	"base-api-auth/data/model"
	"context"
)

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (s *authService) FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.User, error) {
	user, err := s.authRepository.FindByUsernameOrEmail(ctx, usernameOrEmail)
	if err != nil {
		return nil, err
	}

	return user, nil
}
