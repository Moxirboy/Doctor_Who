package usecase

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/pkg/sms"
	"DoctorWho/internal/repository"
	"errors"
)

type usecase struct {
	repo repository.Repo
	f    domain.Factory
	bot  Bot.Bot
}
type Usecase interface {
	RegisterDoctor(doctor *domain.NewUser) (int, error)
	RegisterUser(user *domain.NewUser) (int, error)
	Login(Number string) (bool, int, error)
	Verify(id string, code string) (bool, error)
	GetAll() (User []dto.User)
	FillInfo(user dto.UserInfo) (int, error)
	GetUserInfo(userId int)  (user dto.UserInfo,err error)
	UpdateInfo( user dto.UserInfo) (id int, err error)
	GetProgress(id int) (prog dto.Progress, err error)
}

func NewUserUsecase(repo repository.Repo, bot Bot.Bot) Usecase {
	return &usecase{repo: repo, bot: bot}
}
func (u usecase) RegisterUser(newUser *domain.NewUser) (int, error) {
	exist, err := u.repo.Exist(newUser.PhoneNumber)
	if errors.Is(err, domain.ErrPhoneNumberExist) || exist {
		u.bot.SendErrorNotification(err)
		return 0, domain.ErrPhoneNumberExist
	}
	code := sms.GenerateVerificationCode()
	u.bot.SendNotification(code)
	err = sms.SendEmail(newUser.PhoneNumber, code)
	u.bot.SendNotification(newUser.PhoneNumber)
	if err != nil {
		u.bot.SendErrorNotification(err)
	}
	user := u.f.CreateUser(newUser)
	id, err := u.repo.Register(*user)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	err = u.repo.CreateVerificationCode(id, code)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err

	}
	return id, nil
}
func (u usecase) RegisterDoctor(newUser *domain.NewUser) (int, error) {
	exist, err := u.repo.Exist(newUser.PhoneNumber)
	if errors.Is(err, domain.ErrPhoneNumberExist) || exist {
		u.bot.SendErrorNotification(err)
		return 0, domain.ErrPhoneNumberExist
	}
	doctor := u.f.CreateDoctor(newUser)
	return u.repo.Register(*doctor)
}
func (u usecase) Verify(id string, code string) (bool, error) {
	DatabaseCode, err := u.repo.GetVerificationCode(id)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}
	match := DatabaseCode == code
	if !match {
		u.bot.SendErrorNotification(err)
		return false, err
	}
	_, err = u.repo.UpdateIsUsed(id, code)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}
	return match, nil
}

func (u usecase) Login(Number string) (bool, int, error) {
	exist, err := u.repo.Exist(Number)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, 0, err
	}
	id, err := u.repo.GetByEmail(Number)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, 0, err
	}
	return exist, id, nil
}
func (u usecase) GetAll() (User []dto.User) {
	return u.repo.GetAll()
}
func (u usecase) FillInfo(user dto.UserInfo) (int, error) {
	userInfo := u.f.ParseModelToUserInfo(user)
	id, err := u.repo.CreateInfo(*userInfo)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) UpdateInfo( user dto.UserInfo) (id int, err error){
	userInfo:=u.f.ParseModelToUserInfo(user)
	if !Validator(userInfo.Name){
		id ,err=u.repo.UpdateName(*userInfo)
		if err!=nil{
			u.bot.SendErrorNotification(err)
			return 0,domain.Err("Coudn`t update")
		}
	}
	if !Validator(userInfo.Height){
		id ,err=u.repo.UpdateHeight(*userInfo)
		if err!=nil{
			u.bot.SendErrorNotification(err)
			return 0,domain.Err("Coudn`t update")
		}
	}
	if !Validator(userInfo.Age){
		id ,err=u.repo.UpdateAge(*userInfo)
		if err!=nil{
			u.bot.SendErrorNotification(err)
			return 0,domain.Err("Coudn`t update")
		}
	}
	if !Validator(userInfo.Waist){
		id ,err=u.repo.UpdateWaist(*userInfo)
		if err!=nil{
			u.bot.SendErrorNotification(err)
			return 0,domain.Err("Coudn`t update")
		}
	}
	if !Validator(userInfo.Weigh){
		id ,err=u.repo.UpdateWeigh(*userInfo)
		if err!=nil{
			u.bot.SendErrorNotification(err)
			return 0,domain.Err("Coudn`t update")
		}
	}
	return id, nil
}
func (u usecase) GetUserInfo(userId int)  (user dto.UserInfo,err error){
	userInfo,err:=u.repo.GetUserInfo(userId)
	if err!=nil{
		u.bot.SendErrorNotification(err)
		return user,domain.ErrCouldNotScan
	}
	user=*u.f.ParseUserInfoToModel(userInfo)
	return user,nil
}
func (u usecase) GetProgress(id int) (prog dto.Progress, err error) {
	return prog, err
}
