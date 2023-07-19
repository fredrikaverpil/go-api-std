package models

import "errors"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
}

func (u *User) Validate() error {
	if u.ID <= 0 {
		return errors.New("invalid user id")
	}

	return nil
}
