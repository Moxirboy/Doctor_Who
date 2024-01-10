package repository

import "DoctorWho/internal/domain"

func (r repo) CreateExercise(exercise domain.Exercise) (id int, err error) {
	query := `
		insert into exercise(program_id,name,info,link) values($1,$2,$3,$4) returning id
`
	err = r.db.QueryRow(query, exercise.ProgramId, exercise.Name, exercise.Info, exercise.Link).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) GetAllExercise() (exe []domain.Exercise, err error) {
	query := `
	select id, program_id,name,info,link from exercise
	`

	rows, err := r.db.Query(query)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var exercise domain.Exercise
		err = rows.Scan(
			&exercise.Id,
			&exercise.ProgramId,
			&exercise.Name,
			&exercise.Info,
			&exercise.Link,
		)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotScan
		}
		exe = append(exe, exercise)
	}

	return exe, nil
}
