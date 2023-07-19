package stores

import "github.com/fredrikaverpil/go-api-std/models"

type Store interface {
	GetUser(id int) (models.User, error)
}
