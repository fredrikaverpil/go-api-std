package services

import (
	"github.com/fredrikaverpil/go-api-std/internal/domain"
	"github.com/fredrikaverpil/go-api-std/internal/models"
	"github.com/fredrikaverpil/go-api-std/internal/stores"
)

type UserService struct {
	store stores.Store
}

func NewService(store stores.Store) *UserService {
	service := UserService{
		store: store,
	}
	return &service
}

func (s *UserService) CreateUser(username string, password string) (models.User, error) {
	preExistingUser, err := s.store.GetUserByUsername(username)
	if err != nil {
		if e, ok := err.(*domain.Error); ok && e.Code == domain.ErrNotFound {
			// expected, username should not exist here, or it is already taken
		} else {
			// any other error is an actual problem
			return models.User{}, err
		}
	}

	if preExistingUser.ID != 0 {
		return models.User{}, domain.ConflictError("username already exists")
	}

	user, error := s.store.CreateUser(username, password)
	if error != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (models.User, error) {
	if id <= 0 {
		m := "Not found: record must have id >= 1"
		return models.User{}, domain.NotFoundError(m)
	}

	user, err := s.store.GetUser(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
