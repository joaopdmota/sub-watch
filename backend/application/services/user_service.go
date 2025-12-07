package services

import (
	"context"
	"sub-watch/application/domain"
	"sub-watch/infra/repositories"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
