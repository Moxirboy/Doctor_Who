package domain

import (
	"testDeployment/internal/delivery/dto"
	"time"
)

type Factory struct {
}

func (f Factory) CreateUser(newUser *NewUser) *User {
	return &User{
		phone_number: newUser.PhoneNumber,
		password: newUser.Password,
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
func (f Factory) ParseModelToUserInfo(u dto.UserInfo) *UserInfo {
	return &UserInfo{
		Id: u.Id,
		Name:      u.Name,
		Weigh:     u.Weigh,
		Height:    u.Height,
		Age:       u.Age,
		Waist:     u.Waist,
		Gender: u.Gender,
		UpdatedAt: time.Now(),
	}
}
func (f Factory) ParseUserInfoToModel(u UserInfo)  *dto.UserInfo{
	return &dto.UserInfo{
		Name:      u.Name,
		Weigh:     u.Weigh,
		Height:    u.Height,
		Age:       u.Age,
		Waist:     u.Waist,
		Gender: u.Gender,
	}
}
