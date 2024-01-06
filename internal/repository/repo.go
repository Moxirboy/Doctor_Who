package repository

import (
	"DoctorWho/internal/common/const"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
	"database/sql"
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
	db  *sql.DB
	f   domain.Factory
	Bot Bot.Bot
}
type Repo interface {
	Register(user domain.User) (int, error)
	Exist(email string) (bool, error)
	GetByEmail(email string) (id int, err error)
	GetAll() []dto.User
	UpdatePhoneNumber(number string) (int, error)
	CreateInfo(user domain.UserInfo) (int, error)
	GetUserInfo(userId int) (domain.UserInfo,error)
	UpdateInfo(user domain.UserInfo) (id int,err error)
	UpdateName(user domain.UserInfo) (id int,err error)
	UpdateWeigh(user domain.UserInfo) (id int,err error)
	UpdateHeight(user domain.UserInfo) (id int,err error)
	UpdateAge(user domain.UserInfo) (id int,err error)
	UpdateWaist(user domain.UserInfo) (id int,err error)
	CreateProgram(ageUp, ageDown int, bmi float64, programType _const.ProgramType, proType _const.ProType) (id int, err error)
	UpdateIsUsed(userId string, code string) (id int, err error)
	UpdateVerified(userId string) (id int, err error)
	GetVerificationCode(id string) (code string, err error)
	CreateVerificationCode(id int, code string) (err error)
}

func NewRepo(db *sql.DB, bot Bot.Bot) Repo {
	return &repo{db: db,
		Bot: bot,
	}
}
func (r repo) Register(user domain.User) (id int, err error) {
	query := `
	insert into users (email,role,created_at,updated_at,deleted_at) values($1,$2,$3,$4,$5) returning id
`
	row := r.db.QueryRow(query, user.Phone_number(), user.Role(), user.Created_at(), user.Updated_at(), user.Deleted_at())
	if err := row.Scan(&id); err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (r repo) Exist(email string) (exist bool, err error) {

	query := `
Select Exists (
			SELECT 1
			FROM users
			WHERE email = $1)
		
`
	err = r.db.QueryRow(query, email).Scan(&exist)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return false, domain.ErrCouldNotScan
	}

	return exist, nil
}
func (r repo) GetByEmail(email string) (id int, err error) {
	query := `
		select id from users where email=$1
`
	err = r.db.QueryRow(query, email).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (r repo) GetAll() (User []dto.User) {
	var user userDB
	query := `
		select * from users
`
	rows, err := r.db.Query(query)
	if err != nil {
		r.Bot.SendErrorNotification(err)
	}
	for rows.Next() {

		err := rows.Scan(&user.id, &user.phone_number, &user.role, &user.created_at, &user.updated_at, &user.deleted_at)
		if err != nil {
			r.Bot.SendErrorNotification(err)
		}

		User = append(User, r.f.ParseModelToDomain(user.id, user.phone_number, user.role, user.created_at, user.updated_at, user.deleted_at))
	}
	return User
}
func (r repo) UpdatePhoneNumber(number string) (id int, err error) {
	return 0, err
}

func (r repo) CreateInfo(user domain.UserInfo) (id int, err error) {
	query := `
	insert into  user_info (user_id,name,weigh,height,age,waist,updated_at) values($1,$2,$3,$4,$5,$6,$7) RETURNING id
`
	row := r.db.QueryRow(query, user.Id,user.Name, user.Weigh, user.Height, user.Age, user.Waist, user.UpdatedAt)
	if err = row.Scan(&id); err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}
func (r repo )GetUserInfo(userId int) (user domain.UserInfo,err error){
	query:=`
	select name,weigh,height,age,waist from user_info where user_id=$1
	`
	err=r.db.QueryRow(query,userId).Scan(
		&user.Name,
		&user.Weigh,
		&user.Height,
		&user.Age,
		&user.Waist,
	)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return user,domain.ErrCouldNotScan
	}
	return user,nil
}



func (r repo) UpdateInfo(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set name=$2,weigh=$3,height=$4,age=$5,waist=$6,updated_at=$7 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Name, user.Weigh, user.Height, user.Age, user.Waist, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
}
func (r repo) UpdateName(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set name=$2,updated_at=$3 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Name, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
}
func (r repo) UpdateWeigh(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set weigh=$2,updated_at=$3 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Weigh, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
}
func (r repo) UpdateHeight(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set height=$2,updated_at=$3 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Height, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
	
}
func (r repo) UpdateAge(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set age=$2,updated_at=$3 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Age, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
}
func (r repo) UpdateWaist(user domain.UserInfo) (id int,err error){
	query:=`
	update user_info set waist=$2,updated_at=$3 where user_id=$1 returning id
	`
	err =r.db.QueryRow(query,user.Id,user.Waist, user.UpdatedAt).Scan(&id)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id ,nil
}
func (r repo ) GetDoneExercise(){}














func (r repo) CreateProgram(ageUp, ageDown int, bmi float64, programType _const.ProgramType, proType _const.ProType) (id int, err error) {
	query := `
	insert into programs (type,ageUp,ageDown,bmi,pro_type) values ($1,$2,$3,$4,$5) returning id
`
	err = r.db.QueryRow(query, programType, ageUp, ageDown, bmi, proType).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
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
	rows, err := r.db.Query(query, age, _const.StressWork, _const.Personal)
	if err!=nil{
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			r.Bot.SendErrorNotification(err)
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
	rows, err := r.db.Query(query, age, _const.StressWork, _const.Recommended)
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
func (r repo) GetProgramForWeightLoss(age int, bmi float64) (int, error) {
	var ids []int
	query := `
		select id from programs where ageUp>$1 and ageDown<$1 and bmiUp>$2 and bmiDown<$2 and type=$3 and pro_type=$4
`
	rows, err := r.db.Query(query, age, bmi, _const.WeightLoss, _const.Personal)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotRetrieveFromDataBase
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return 0, domain.ErrCouldNotScan
		}
		ids = append(ids, id)
	}
	id := r.Random(ids)
	return id, nil
}
func (r repo) GetRecommendedProgramForWeightLoss(age int, bmi float64) (ids []int, err error) {

	query := `
		select id from programs where ageUp>$1 and ageDown<$1 and bmiUp>$2 and bmiDown<$2 and type=$3 and pro_type=$4
`
	rows, err := r.db.Query(query, age, bmi, _const.WeightLoss, _const.Personal)
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
func (r repo) CreateProgramChosen(userId, programId int) (id int, err error) {
	query := `
		insert into program_chosen(program_id,user_id) values($1,$2) returning id
`
	err = r.db.QueryRow(query).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}
func (r repo) CreateExercise(name, info string, programId int) (id int, err error) {
	query := `
		insert into exercise(program_id,name,info) values($1,$2,$3) returning id
`
	err = r.db.QueryRow(query, programId, name, info).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, domain.ErrCouldNotScan
	}
	return id, nil
}

//func (r repo) GetExercises()

func (r repo) GetAllExercise() {
	// TODO implement
}
func (r repo) GetExerciseByProgram() {
	//Todo implement
}
func (r repo) UpdateExercise() {
	// TODO implement
}
func (r repo) DeleteExercise() {
	// TODO impolement
}
func (r repo) CreateExerciseChose() {
	//Todo implement
}
func (r repo) GetAllExerciseChoosen() {
	//Todo implement
}

func (r repo) GetExerciseChoosenByUserId() {
	//todo implement
}
func (r repo) GetExerciseChoosenByProgramId() {
	//todo implement
}
func (r repo) GetDoneExerciseChoosenByUserID() {
	//todo implement
}
func (r repo) UpdateExerciseChoosen() {

}
func (r repo) CreateVerificationCode(id int, code string) (err error) {
	query := `
		insert into verify_emails(user_id,secret_code,is_used) values($1,$2,$3)
`
	_, err = r.db.Exec(query, id, code, false)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return err
	}
	return nil

}

func (r repo) GetVerificationCode(id string) (code string, err error) {
	query := `
		select secret_code from verify_emails where user_id=$1 and is_used=$2
`
	err = r.db.QueryRow(query, id, false).Scan(&code)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return "", domain.ErrCouldNotScan
	}

	return code, nil
}
func (r repo) UpdateIsUsed(userId string, code string) (id int, err error) {
	query := `
		update verify_emails set is_used=$1 where user_id=$2 and secret_code=$3 returning id
	`
	err = r.db.QueryRow(query, true, userId, code).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err

	}
	return id, nil
}
func (r repo) UpdateVerified(userId string) (id int, err error) {
	query := `
		Update users set is_email_verified=$1 where user_id=$2 returning id
`
	err = r.db.QueryRow(query, true, userId).Scan(&id)
	if err != nil {
		r.Bot.SendErrorNotification(err)
		return 0, err
	}
	return id, nil
}

//func (r repo ) GetRecommended

func (r repo) Random(ids []int) int {
	// TODO implement
	return 0
}
