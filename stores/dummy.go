// these are all dummy store functions for testing purposes

package stores

import (
	"errors"

	"github.com/fredrikaverpil/go-api-std/models"
)

type DummyStore struct{}

func (s *DummyStore) GetUser(id int) (models.User, error) {
	if id == 1 {
		user := models.User{
			ID:        id,
			FirstName: "John",
		}

		return user, nil
	} else {
		return models.User{}, errors.New("no user found")
	}
}
