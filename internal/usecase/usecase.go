package usecase

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/repository"
	"errors"
)

type usecase struct {
	repo repository.Repo
	f    domain.Factory
}
type Usecase interface {
	RegisterDoctor(doctor *domain.NewUser) (int, error)
	RegisterUser(user *domain.NewUser) (int, error)
	Login(Number string) (bool, error)
	GetAll() (User []dto.User)
	FillInfo(user dto.UserInfo) (int, error)
}

func NewUserUsecase(repo repository.Repo) Usecase {
	return &usecase{repo: repo}
}
func (u usecase) RegisterUser(newUser *domain.NewUser) (int, error) {
	exist, err := u.repo.Exist(newUser.PhoneNumber)
	if errors.Is(err, domain.ErrPhoneNumberExist) || exist {
		return 0, domain.ErrPhoneNumberExist
	}
	user := u.f.CreateUser(newUser)
	return u.repo.Register(*user)
}
func (u usecase) RegisterDoctor(newUser *domain.NewUser) (int, error) {
	exist, err := u.repo.Exist(newUser.PhoneNumber)
	if errors.Is(err, domain.ErrPhoneNumberExist) || exist {
		return 0, domain.ErrPhoneNumberExist
	}
	doctor := u.f.CreateDoctor(newUser)
	return u.repo.Register(*doctor)
}
func (u usecase) Login(Number string) (bool, error) {
	return u.repo.Exist(Number)
}
func (u usecase) GetAll() (User []dto.User) {
	return u.repo.GetAll()
}
func (u usecase) FillInfo(user dto.UserInfo) (int, error) {
	return u.repo.UpdateInfo(user)
}
