package services

import (
	"github.com/sergiocltn/api-go-std/models"
	repositories "github.com/sergiocltn/api-go-std/repository"
)

type IUserService interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(*models.User) (int, error)
}

type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	return s.repo.GetUserByIDQuery(id)
}

func (s *UserService) CreateUser(user *models.User) (int, error) {
	//TODO
	return s.repo.CreateUserQuery(*user)
}
