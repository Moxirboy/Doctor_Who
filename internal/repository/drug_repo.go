package repository

import (
	"strings"
	"testDeployment/internal/domain"
)
func ( r repo) GetAllDrug()(drugs []domain.Drug,err error){
	query:=`
	select * from drug 
	`
	rows,err:=r.db.Query(query)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return nil, err
	}
	for rows.Next() {
		var drug domain.Drug
		err = rows.Scan(
			&drug.Id,
			&drug.Name,
			&drug.Description,
			&drug.Manufacturer,
			&drug.Receipt,
		)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotRetrieveFromDataBase
		}
		drugs = append(drugs, drug)
	}
	return drugs,nil
}
func (r repo) InsertDrug(drug domain.Drug) (id int, err error) {
	query := `
	insert into drug(name,description,manufacturer,reciept) values($1,$2,$3,$4) returning id
	`
	err = r.db.QueryRow(query, drug.Name, drug.Description, drug.Manufacturer, drug.Receipt).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}

func (r repo) GetDrugByName(name string) (drugs []domain.Drug, err error) {
	query := `
    SELECT * FROM drug WHERE LOWER(name) LIKE LOWER($1)
`

	rows, err := r.db.Query(query, "%"+strings.ToLower(name)+"%")
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return nil, err
	}
	for rows.Next() {
		var drug domain.Drug
		err = rows.Scan(
			&drug.Id,
			&drug.Name,
			&drug.Description,
			&drug.Manufacturer,
			&drug.Receipt,
		)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return nil, domain.ErrCouldNotRetrieveFromDataBase
		}
		drugs = append(drugs, drug)
	}
	return drugs, nil
}
func (r repo) GetDrugById(id string) (drug domain.Drug, err error) {
	query := `
    SELECT * FROM drug WHERE id=$1
`

	err = r.db.QueryRow(query, id).Scan(
		&drug.Id,
		&drug.Name,
		&drug.Description,
		&drug.Manufacturer,
		&drug.Receipt,
	)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return drug, err
	}

	return drug, nil
}
