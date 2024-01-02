package domain

import "time"

type Factory struct {
}

func (f Factory) CreateUser(newUser *NewUser) *User {
	return &User{
		phone_number: newUser.PhoneNumber,
		password:     newUser.Password,
		role:         "user",
		created_at:   time.Now().UTC(),
		updated_at:   time.Now().UTC(),
		deleted_at:   nil,
	}
}
func (f Factory) CreateDoctor(newUser *NewUser) *User {
	return &User{
		phone_number: newUser.PhoneNumber,
		password:     newUser.Password,
		role:         "doctor",
		created_at:   time.Now().UTC(),
		updated_at:   time.Now().UTC(),
		deleted_at:   nil,
	}
}
