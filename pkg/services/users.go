package services

import (
	"github.com/fredrikaverpil/go-api-std/pkg/lib"
	"github.com/fredrikaverpil/go-api-std/pkg/models"
	"github.com/fredrikaverpil/go-api-std/pkg/stores"
)

func CreateUser(store stores.Store, username string, password string) (models.User, error) {
	preExistingUser, err := store.GetUserByUsername(username)
	if err != nil {
		ierr := err.(*lib.CustomError)
		switch ierr.Code {
		case lib.ErrNotFound:
			// expected, username should not exist here, or it is already taken
			break
		default:
			// any other error is an actual problem
			return models.User{}, err
		}
	}
	if preExistingUser.ID != 0 {
		return models.User{}, lib.ConflictError("username already exists")
	}

	user, error := store.CreateUser(username, password)
	if error != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUser(store stores.Store, id int) (models.User, error) {
	if id <= 0 {
		m := "Not found: record must have id >= 1"
		return models.User{}, lib.NotFoundError(m)
	}

	user, err := store.GetUser(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
