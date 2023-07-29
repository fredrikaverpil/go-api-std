package stores

import "github.com/fredrikaverpil/go-api-std/pkgs/models"

type Store interface {
	CreateUser(username string, password string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}
