package usecase

import "testDeployment/internal/domain"

func (u usecase) CreateExercise(exercise domain.Exercise) (id int, err error) {
	id, err = u.repo.CreateExercise(exercise)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) GetAllExercise() (exe []domain.Exercise, err error) {
	exe, err = u.repo.GetAllExercise()
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, err
	}
	return exe, nil
}
