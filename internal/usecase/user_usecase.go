package usecase

import (
	"DoctorWho/internal/domain"
	"DoctorWho/internal/repository"
)

type usecase struct {
	repo repository.Repo
	f    domain.Factory
}
type Usecase interface {
	RegisterDoctor(doctor *domain.NewUser) (int, error)
	RegisterUser(user *domain.NewUser) (int, error)
	Login(Number string, pass string) (bool, error)
}

func NewUserUsecase(repo repository.Repo) Usecase {
	return &usecase{repo: repo}
}
func (u usecase) RegisterUser(newUser *domain.NewUser) (int, error) {
	user := u.f.CreateUser(newUser)
	return u.repo.Register(*user)
}
func (u usecase) RegisterDoctor(newUser *domain.NewUser) (int, error) {
	doctor := u.f.CreateDoctor(newUser)
	return u.repo.Register(*doctor)
}
func (u usecase) Login(Number string, pass string) (bool, error) {
	return u.repo.Exist(Number, pass)
}
