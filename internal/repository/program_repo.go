package repository

import (
	"log"
	_const "DoctorWho/internal/common/const"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
)

func (r repo) CreateProgram(pro domain.Program, proType _const.ProType) (id int, err error) {
	query := `
	insert into programs (type,ageUp,ageDown,bmiUp,bmiDown,pro_type) values ($1,$2,$3,$4,$5,$6) returning id
`
	err = r.db.QueryRow(query, pro.Type, pro.AgeUp, pro.AgeDown, pro.BMIUp, pro.BMIDown, proType).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotCreateProgram
	}
	return id, nil
}

func (r repo) GetAllPrograms() (pros []domain.Program, err error) {
	query := `
	select  id, ageUp,ageDown,bmiUp,bmiDown,type from programs
	`
	rows, err := r.db.Query(query)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotCreateProgram
	}
	for rows.Next() {
		var pro domain.Program
		err = rows.Scan(
			&pro.ID,
			&pro.AgeUp,
			&pro.AgeDown,
			&pro.BMIUp,
			&pro.BMIDown,
			&pro.Type,
		)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return pros, err
		}
		pros = append(pros, pro)
	}
	log.Println("pros:",pros)
	return pros, nil
}

func (r repo) GetRecommendedProgram(age int, bmi float64, tip dto.ProgramType) (ids []int, err error) {

	query := `
		select id from programs where  CAST(ageUp AS INTEGER)>$1 and  CAST(ageDown AS INTEGER)<$1 and CAST(bmiUp AS DECIMAL)>$2 and CAST(bmiDown AS DECIMAL) <$2 and type=$3 and pro_type=$4
`
	rows, err := r.db.Query(query, age, bmi, tip, _const.Personal)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotScan
		}
		ids = append(ids, id)
	}

	return ids, nil
}
