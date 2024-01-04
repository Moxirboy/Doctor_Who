package repository

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"database/sql"
	"log"
	"time"
)

type userDB struct {
	id           int
	phone_number string
	role         string
	created_at   time.Time
	updated_at   time.Time
	deleted_at   *time.Time
}
type repo struct {
	db *sql.DB
	f  domain.Factory
}
type Repo interface {
	Register(user domain.User) (int, error)
	Exist(Number string) (bool, error)
	GetAll() []dto.User
	UpdatePhoneNumber(number string) (int, error)
	UpdateInfo(user dto.UserInfo) (int, error)
	CreateProgram(ageUp, ageDown int, bmi float64, programType domain.ProgramType, proType domain.ProType) (id int, err error) }

func NewRepo(db *sql.DB) Repo {
	return &repo{db: db}
}
func (r repo) Register(user domain.User) (id int, err error) {
	query := `
	insert into users (phone_number,password,role,created_at,updated_at,deleted_at) values($1,$2,$3,$4,$5,$6) returning id
`
	row := r.db.QueryRow(query, user.Phone_number(), user.Role(), user.Created_at(), user.Updated_at(), user.Deleted_at())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r repo) Exist(Number string) (exist bool, err error) {

	query := `
Select Exists (
			SELECT 1
			FROM users
			WHERE phone_number = $1)
		
`
	err = r.db.QueryRow(query, Number).Scan(&exist)
	if err != nil {
		return false, domain.ErrCouldNotScan
	}

	return exist, nil
}
func (r repo) GetAll() (User []dto.User) {
	var user userDB
	query := `
		select * from users
`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {

		err := rows.Scan(&user.id, &user.phone_number, &user.role, &user.created_at, &user.updated_at, &user.deleted_at)
		if err != nil {
			log.Println(err)
		}

		User = append(User, r.f.ParseModelToDomain(user.id, user.phone_number, user.role, user.created_at, user.updated_at, user.deleted_at))
	}
	return User
}
func (r repo) UpdatePhoneNumber(number string) (id int, err error) {
	return 0, err
}
func (r repo) UpdateInfo(user dto.UserInfo) (id int, err error) {
	query := `
	update user_info set user_id=$1,name=$2,weigh=$3,height=$4,age=$5,waist=$6 RETURNING id
`
	row := r.db.QueryRow(query, user.Name, user.Weigh, user.Height, user.Age, user.Waist)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r repo) CreateProgram(ageUp, ageDown int, bmi float64, programType domain.ProgramType, proType domain.ProType) (id int, err error) {
	query := `
	insert into programs (type,ageUp,ageDown,bmi,pro_type) values ($1,$2,$3,$4,$5) returning id
`
	err = r.db.QueryRow(query, programType, ageUp, ageDown, bmi, proType).Scan(&id)
	if err != nil {
		return 0, domain.ErrCouldNotCreateProgram
	}
	return id, nil
}

// func (r repo ) GetAllPrograms()[]
func (r repo) GetProgramByAge(age int) (int, error) {
	var ids []int
	query := `
	select id from programs where  ageUp>$1 and ageDown<$1 and type=$2 and pro_type=$3
`
	rows, err := r.db.Query(query, age, domain.StressWork,domain.Personal)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return 0, domain.ErrCouldNotScan
		}
		ids = append(ids, id)
	}
	return r.Random(ids), nil
}

func (r repo) GetRecommendedProgramByAge(age int) ([]int, error) {
	var ids []int
	query := `
	select id from programs  where  ageUp>$1 and ageDown<$1 and type=$2 and pro_type=$3
`
	rows, err := r.db.Query(query, age, domain.StressWork, domain.Recommended)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, domain.ErrCouldNotScan
		}
		ids = append(ids, id)
	}
	return ids, nil
}
func (r repo ) GetProgramForWeightLoss(age int, bmi float64) (int,error){
	var ids []int
	query:=`
		select id from programs where ageUp>$1 and ageDown<$1 and bmiUp>$2 and bmiDown<$2 and type=$3 and pro_type=$4
`
	rows,err:=r.db.Query(query,age,bmi,domain.WeightLoss,domain.Personal)
	if err!=nil{
		return 0, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next(){
		var id int
		err=rows.Scan(&id)
		if err!=nil{
			return 0, domain.ErrCouldNotScan
		}
		ids=append(ids,id)
	}
	id :=r.Random(ids)
	return id,nil
}
func (r repo ) GetRecommendedProgramForWeightLoss(age int, bmi float64) (ids []int,err error){

	query:=`
		select id from programs where ageUp>$1 and ageDown<$1 and bmiUp>$2 and bmiDown<$2 and type=$3 and pro_type=$4
`
	rows,err :=r.db.Query(query,age,bmi,domain.WeightLoss,domain.Personal)
	if err!=nil{
		return nil, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next(){
		var id int
		err=rows.Scan(&id)
		if err!=nil{
			return nil, domain.ErrCouldNotScan
		}
		ids=append(ids,id)
	}

	return ids,nil
}
func (r repo) CreateProgramChosen(userId,programId int) (id int ,err error){
	query:=`
		insert into program_chosen(program_id,user_id) values($1,$2) returning id
`
	err=r.db.QueryRow(query).Scan(&id)
	if err!=nil{
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
CREATE TABLE exercise (
id SERIAL PRIMARY KEY,
program_id INT,
name VARCHAR(255),
info varchar(255),
CONSTRAINT fk_exercise_program FOREIGN KEY (program_id) REFERENCES programs(id)
);

func (r repo ) CreateExercise(name ,info string,programId int) (id int ,err error){
	query:=`
		insert into exercise(program_id,name,info) values($1,$2,$3) returning id
`
	err=r.db.QueryRow(query,programId ,name,info).Scan(&id)
	if err != nil {
		return 0, domain.ErrCouldNotScan
	}
	return id,nil
}
func (r repo) GetAllExercise(){
	// TODO implement
}
func (r repo) GetExerciseByProgram(){
	//Todo implement
}
func (r repo ) UpdateExercise(){
	// TODO implement
}
func (r repo) DeleteExercise(){
	// TODO impolement
}
func (r repo ) CreateExerciseChose(){
	//Todo implement
}
func(r repo) GetAllExerciseChoosen(){
	//Todo implement
}

func (r repo ) GetExerciseChoosenByUserId(){
	//todo implement
}
func (r repo) GetExerciseChoosenByProgramId(){
	//todo implement
}
func (r repo ) GetDoneExerciseChoosenByUserID(){
	//todo implement
}
func ( r repo) UpdateExerciseChoosen(){

}


//func (r repo ) GetRecommended



func (r repo ) Random(ids []int) int{
	// TODO implement
	return 0
}