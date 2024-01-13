package user

import (
	"github.com/sirupsen/logrus"
	"go-gin-mall/internal/driving/repository"
)

type UserDomainService struct {
	Logger *logrus.Logger
	Repo   *repository.UserRepository
}

func NewUserDomainService(logger *logrus.Logger, repo *repository.UserRepository) *UserDomainService {
	return &UserDomainService{
		Repo:   repo,
		Logger: logger,
	}
}
