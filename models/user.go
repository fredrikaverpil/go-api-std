package models

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (u *User) Validate() error {
	if u.ID <= 0 {
		return errors.New("invalid user id")
	}
	if u.Username == "" {
		return errors.New("invalid username")
	}

	return nil
}
