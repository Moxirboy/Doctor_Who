package usecase
import (
	"errors"
	"testDeployment/internal/domain"
)

func (u usecase) RegisterDoctor(newUser *domain.NewUser) (int, error) {
	exist, err := u.repo.Exist(newUser.PhoneNumber)
	if errors.Is(err, domain.ErrPhoneNumberExist) || exist {
		u.bot.SendErrorNotification(err)
		return 0, domain.ErrPhoneNumberExist
	}
	doctor := u.f.CreateDoctor(newUser)
	return u.repo.Register(*doctor)
}
func (u usecase) FillDoctorInfo(info domain.Doctor) (int ,error){
	empty:=isAnyFieldEmpty(info)
	if empty{
		return 0 ,domain.ErrEmptyField
	}
	id,err:=u.repo.FillDoctorInfo(info)
	if err!=nil{
		return 0 ,err
	}
	return id ,err
}
// func (u usecase) ChangeInfo()
// func (u usecase) ChangePhoto(photo string) (id  int ,err error){

// }