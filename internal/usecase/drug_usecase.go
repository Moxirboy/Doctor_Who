package usecase

import "testDeployment/internal/domain"

func (u usecase) CreateDrug(drug domain.Drug) (id int, err error) {
	id, err = u.repo.InsertDrug(drug)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	err = u.repo.CreatePhoto(id, drug.Photo)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (u usecase) GetDrugs(drugS domain.DrugSearch) (drugs []domain.Drug, err error) {
	drugs, err = u.repo.GetDrugByName(drugS.Name)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, err
	}
	return drugs, nil
}

func (u usecase) GetDrug(d domain.DrugSearch) (drug domain.Drug, err error) {
	drug, err = u.repo.GetDrugById(d.Id)
	if err != nil {
		u.bot.SendErrorNotification(err)
		return drug, err
	}
	return drug, nil
}
func (u usecase) GetAllDrug()(drugs []domain.Drug,err error){
	drugs, err = u.repo.GetAllDrug()
	if err != nil {
		u.bot.SendErrorNotification(err)
		return nil, err
	}
	return drugs, nil
}