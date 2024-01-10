package repository

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
)

func (r repo) GetDoneExercise(personal domain.PersonalExercisesDone) (booller []bool, err error) {
	query := `
	SELECT
			ec.done
		FROM
			exercise_chosen ec
		JOIN
			exercise e ON ec.exercise_id = e.id
		JOIN
			programs p ON e.program_id = p.id
		WHERE
			ec.user_id = $1
			AND DATE(ec.created_at) = $2
			AND p.type = $3
			AND p.pro_type = 'personal'
	`
	rows, err := r.db.Query(query, personal.UserId, personal.Date, personal.Typo)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var done bool
		err = rows.Scan(&done)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotScan
		}
		booller = append(booller, done)
	}
	return booller, nil

}

func (r repo) GetPersonalExerciseChoosen(userId int, date string, proType string) (exercises []domain.PersonalExercises, err error) {
	query := `
	SELECT
			ec.id AS chosen_id,
			ec.done,
			e.name AS exercise_name,
			e.info AS exercise_info,
			e.link AS exercise_link
		FROM
			exercise_chosen ec
		JOIN
			exercise e ON ec.exercise_id = e.id
		JOIN
			programs p ON e.program_id = p.id
		WHERE
			ec.user_id = $1
			AND DATE(ec.created_at) = $2
			AND p.type = $3
			AND p.pro_type = 'personal'
	`

	rows, err := r.db.Query(query, userId, date, proType)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var exercise domain.PersonalExercises
		err = rows.Scan(
			&exercise.Id,
			&exercise.Done,
			&exercise.Name,
			&exercise.Info,
			&exercise.Link,
		)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotScan
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}
func (r repo) UpdateDone(mark domain.MarkAsDone) (id int, err error) {
	query := `
		update exercise_chosen set done=$1 where id=$2  returning id
	`
	err = r.db.QueryRow(query, mark.Done, mark.Id).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) GetExercises(programId int) (ids []int, err error) {
	query := `
	select id from exercise where program_id=$1
`
	rows, err := r.db.Query(query, programId)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, err
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}
	return ids, nil
}

func (r repo) CreateExerciseChoosen(id, userId int, tip dto.ProgramType, createdAt string) (err error) {
	query := `
		insert into exercise_chosen(exercise_id,user_id,created_at,type) values($1,$2,$3,$4)
`
	_, err = r.db.Exec(query, id, userId, createdAt, tip)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return err
	}
	return nil
}
