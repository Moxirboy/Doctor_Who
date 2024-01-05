package domain

import (
	"DoctorWho/internal/delivery/dto"
	"time"
)

type Factory struct {
}

func (f Factory) CreateUser(newUser *NewUser) *User {
	return &User{
		phone_number: newUser.PhoneNumber,
		role:         "user",
		created_at:   time.Now().UTC(),
		updated_at:   time.Now().UTC(),
		deleted_at:   nil,
	}
}
func (f Factory) CreateDoctor(newUser *NewUser) *User {
	return &User{
		phone_number: newUser.PhoneNumber,
		role:         "doctor",
		created_at:   time.Now().UTC(),
		updated_at:   time.Now().UTC(),
		deleted_at:   nil,
	}
}
func (f Factory) ParseModelToDomain(id int, phoneNumber string, role string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) dto.User {
	return dto.User{
		Email: phoneNumber,
	}
}
func (f Factory) ParseDomainToModel(u User) {

}
