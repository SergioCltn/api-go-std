package repositories

import (
	"net/http"
	"sync"

	"github.com/sergiocltn/api-go-std/models"
)

type UserLocalStorage struct {
	users map[int]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[int]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(user *models.User) error {
	s.mutex.Lock()
	s.users[user.ID] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(name string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Name == name {
			return user, nil
		}
	}

	return nil, http.ErrLineTooLong
}
