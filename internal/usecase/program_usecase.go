package usecase

import (
	_const "DoctorWho/internal/common/const"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"time"
)

func (u usecase) GetProgramForWeightLoss(userId int) (exercise []domain.PersonalExercises, err error) {
	today := time.Now().Format("2006/01/02")
	exercise, err = u.repo.GetPersonalExerciseChoosen(userId, today, string(dto.WeightLoss))
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	return exercise, nil
}

func (u usecase) GetProgramForStress(userId int) (exercise []domain.PersonalExercises, err error) {
	today := time.Now().Format("2006/01/02")
	exercise, err = u.repo.GetPersonalExerciseChoosen(userId, today, string(dto.StressWork))
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	return exercise, nil
}

func (u usecase) CreateProgram(pro domain.Program) (id int, err error) {
	id, err = u.repo.CreateProgram(pro, _const.Personal)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) GetAllPrograms() (pro []domain.Program, err error) {
	pro, err = u.repo.GetAllPrograms()
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, err
	}
	return pro, nil
}
