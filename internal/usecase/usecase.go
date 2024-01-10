package usecase

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/repository"
)

type usecase struct {
	repo repository.Repo
	f    domain.Factory
	bot  Bot.Bot
}
type Usecase interface {
	RegisterDoctor(doctor *domain.NewUser) (int, error)
	RegisterUser(user *domain.NewUser) (int, error)
	Exist(newUser domain.NewUser) (bool, error)
	Login(user domain.NewUser) (bool, int, error)
	GetAll() (User []dto.User)
	DeleteUser(id int) (err error)
	FillInfo(user dto.UserInfo) (int, error)
	GetUserInfo(userId int) (user dto.UserInfo, err error)
	UpdateInfo(user dto.UserInfo) (id int, err error)
	UpdateIsVerified(userId interface{}) (err error)
	GetProgramForWeightLoss(userId int) (exercise []domain.PersonalExercises, err error)
	GetProgress(personal domain.PersonalExercisesDone) (prog float64, err error)
	MarkAsDone(mark domain.MarkAsDone) (id int, err error)
	GetProgramForStress(userId int) (exercise []domain.PersonalExercises, err error)
	CreateDrug(drug domain.Drug) (id int, err error)
	GetDrugs(drugS domain.DrugSearch) (drugs []domain.Drug, err error)
	GetDrug(d domain.DrugSearch) (drug domain.Drug, err error)
	CreateProgram(pro domain.Program) (id int, err error)
	GetAllPrograms() (pro []domain.Program, err error)
	CreateExercise(exercise domain.Exercise) (id int, err error)
	GetAllExercise() (exe []domain.Exercise, err error)
	AutoExercise(userId int, programType dto.ProgramType) (bool, string, error)
	autoExerciseCreate(userId int, programType dto.ProgramType, date string) (bool, error)
	createExerciseChoices(exercises []int, userId int, programType dto.ProgramType, date string) error
	GetAllDrug()(drugs []domain.Drug,err error)
}

func NewUserUsecase(repo repository.Repo, bot Bot.Bot) Usecase {
	return &usecase{repo: repo, bot: bot}
}
