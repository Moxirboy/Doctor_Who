package usecase

import (
	"errors"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
)

func (u usecase) FillInfo(user dto.UserInfo) (int, error) {
	userInfo := u.f.ParseModelToUserInfo(user)
	id, err := u.repo.CreateInfo(*userInfo)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) UpdateInfo(user dto.UserInfo) (id int, err error) {
	userInfo := u.f.ParseModelToUserInfo(user)
	if !Validator(userInfo.Name) {

		id, err = u.repo.UpdateName(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Name")
		}
	}
	if !Validator(userInfo.Height) {

		id, err = u.repo.UpdateHeight(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Height")
		}
	}
	if !ValidatorAge(userInfo.Age) {
		id, err = u.repo.UpdateAge(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Age")
		}
	}
	if !Validator(userInfo.Waist) {
		id, err = u.repo.UpdateWaist(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Waist")
		}
	}
	if !Validator(userInfo.Weigh) {
		id, err = u.repo.UpdateWeigh(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Weigh")
		}
	}
	if !Validator(userInfo.Gender) {
		id, err = u.repo.UpdateGender(*userInfo)
		if err != nil {
			u.bot.SendErrorNotification(err)
			return 0, domain.Err("Coudn`t update Gender")
		}
	}
	return id, nil
}
func (u usecase) GetUserInfo(userId int) (user dto.UserInfo, err error) {
	exist, err := u.repo.ExistUserInfo(userId)
	if err != nil || errors.Is(err, domain.ErrCouldNotScan) {
		u.bot.SendErrorNotification(err)
		return user, errors.New("no info")
	}
	if !exist {
		return user, errors.New("no info")
	}
	userInfo, err := u.repo.GetUserInfo(userId)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return user, domain.ErrCouldNotScan
	}
	user = *u.f.ParseUserInfoToModel(userInfo)
	return user, nil
}
