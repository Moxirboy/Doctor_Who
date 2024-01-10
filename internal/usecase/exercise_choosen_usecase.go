package usecase

import (
	"errors"
	"testDeployment/internal/delivery/dto"
	"testDeployment/internal/domain"
	"time"
)

// createExerciseChoices creates exercise choices for the given exercises.
func (u usecase) createExerciseChoices(exercises []int, userId int, programType dto.ProgramType, date string) error {
	for _, exercise := range exercises {
		err := u.repo.CreateExerciseChoosen(exercise, userId, programType, date)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u usecase) AutoExercise(userId int, programType dto.ProgramType) (bool, string, error) {
	date := time.Now().Format("2006-01-02")

	exist, err := u.autoExerciseCreate(userId, programType, date)
	if err != nil {
		return false, "", err
	}

	return exist, date, nil
}

// AutoExerciseCreate generates exercise programs based on the specified program type.
func (u usecase) autoExerciseCreate(userId int, programType dto.ProgramType, date string) (bool, error) {
	exist, err := u.repo.ExistUserInfo(userId)
	if err != nil || errors.Is(err, domain.ErrCouldNotScan) {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	if !exist {
		return false, nil
	}

	userInfo, err := u.repo.GetUserInfo(userId)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	bmi := calculateBMI(userInfo.Weigh, userInfo.Height)
	programs, err := u.repo.GetRecommendedProgram(userInfo.Age, bmi, programType)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	programId, err := getRandomElement(programs)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	exerciseIds, err := u.repo.GetExercises(programId)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	randomExercises, err := getRandomExercises(exerciseIds)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	err = u.createExerciseChoices(randomExercises, userId, programType, date)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return false, err
	}

	return true, nil
}

func (u usecase) MarkAsDone(mark domain.MarkAsDone) (id int, err error) {
	id, err = u.repo.UpdateDone(mark)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) GetProgress(personal domain.PersonalExercisesDone) (prog float64, err error) {
	booller, err := u.repo.GetDoneExercise(personal)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0.0, err
	}
	prog = calculatePercentage(booller)
	return prog, nil
}
