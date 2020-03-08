package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)


type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type Service interface {
	CreateUser(ctx context.Context, email, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository,
		logger,
	}
}

func (s *service) CreateUser(ctx context.Context, email, password string) (string, error) {
	logger := log.With(s.logger, "Service", "Method", "CreateUser")
	user := new(User)
	user.ID = uuid.New().String()
	user.Email = email
	user.Password = password
	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("CreateUser", user.ID)
	return user.ID, nil
}

func (s *service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "Service", "Method", "GetUser")
	email, err := s.repository.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("GetUser", id)
	return email, nil
}
