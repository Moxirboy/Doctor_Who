package repository

import (
	"database/sql"
	"DoctorWho/internal/common/const"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
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
	GetByEmail(email string) (id int, password string, err error)
	GetAll() []dto.User
	UpdatePhoneNumber(number string) (int, error)
	UpdateIsActive(id int, deleteAt time.Time) (err error)
	UpdateUserInfoDeleted(id int, deleteAt time.Time) (err error)
	CreateInfo(user domain.UserInfo) (int, error)
	GetUserInfo(userId int) (domain.UserInfo, error)
	ExistUserInfo(userId int) (bool, error)
	UpdateInfo(user domain.UserInfo) (id int, err error)
	UpdateName(user domain.UserInfo) (id int, err error)
	UpdateWeigh(user domain.UserInfo) (id int, err error)
	UpdateHeight(user domain.UserInfo) (id int, err error)
	UpdateAge(user domain.UserInfo) (id int, err error)
	UpdateWaist(user domain.UserInfo) (id int, err error)
	UpdateGender(user domain.UserInfo) (id int, err error)
	CreateProgram(pro domain.Program, proType _const.ProType) (id int, err error)
	UpdateVerified(userId interface{}) (err error)
	GetPersonalExerciseChoosen(userId int, date string, proType string) (exercises []domain.PersonalExercises, err error)
	UpdateDone(mark domain.MarkAsDone) (id int, err error)
	GetDoneExercise(personal domain.PersonalExercisesDone) (booller []bool, err error)
	InsertDrug(drug domain.Drug) (id int, err error)
	CreatePhoto(id int, path []string) (err error)
	GetDrugByName(name string) (drugs []domain.Drug, err error)
	GetDrugById(id string) (drug domain.Drug, err error)
	GetAllPrograms() (pros []domain.Program, err error)
	CreateExercise(exercise domain.Exercise) (id int, err error)
	GetAllExercise() (exe []domain.Exercise, err error)
	GetRecommendedProgram(age int, bmi float64, tip dto.ProgramType) (ids []int, err error)
	GetExercises(programId int) (ids []int, err error)
	CreateExerciseChoosen(id, userId int, tip dto.ProgramType, createdAt string) (err error)
	GetAllDrug()(drugs []domain.Drug,err error)
	FillDoctorInfo(info domain.Doctor) (id int, err error)
}

func NewRepo(db *sql.DB, bot Bot.Bot) Repo {
	return &repo{db: db,
		Bot: bot,
	}
}
