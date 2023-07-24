// these are all dummy store functions for testing purposes

package stores

import (
	"errors"

	"github.com/fredrikaverpil/go-api-std/lib"
	"github.com/fredrikaverpil/go-api-std/models"
)

type DummyDbRecord struct {
	ID             int
	Username       string
	HashedPassword string
}

var userDb = make(map[int]DummyDbRecord)

type DummyStore struct{}

func (s *DummyStore) CreateUser(username string, password string) (models.User, error) {
	nextID := len(userDb) + 1
	hashedPassword, err := lib.HashPassword(password)
	if err != nil {
		return models.User{}, errors.New("could not hash password")
	}

	user := models.User{ID: nextID, Username: username}
	err = user.Validate()
	if err != nil {
		return models.User{}, errors.New("could not validate user")
	}

	userRecord := DummyDbRecord{ID: nextID, Username: username, HashedPassword: hashedPassword}
	userDb[nextID] = userRecord

	return user, nil
}

func (s *DummyStore) GetUserByUsername(username string) (models.User, error) {
	for userId, dbRecord := range userDb {
		if dbRecord.Username == username {
			user, err := s.GetUser(userId)
			if err != nil {
				return models.User{}, errors.New("could not get user")
			}
			return user, nil
		}
	}

	return models.User{}, errors.New("no user found")
}

func (s *DummyStore) GetUser(id int) (models.User, error) {
	userRecord, exists := userDb[id]
	if !exists {
		return models.User{}, errors.New("no user found")
	}

	user := models.User{ID: userRecord.ID, Username: userRecord.Username}
	err := user.Validate()
	if err != nil {
		return models.User{}, errors.New("could not validate user")
	}

	return user, nil
}
