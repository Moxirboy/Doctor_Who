package usecase

import (
	"context"
	"testDeployment/internal/delivery/dto"
	"testDeployment/internal/domain"
	"testDeployment/pkg/Bot"
	"testDeployment/internal/repository"
	"testDeployment/pkg/utils"
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
	CreateDrug(drug domain.Drug) (id int, err error)
	GetDrugs(drugS domain.DrugSearch) (drugs []domain.Drug, err error)
	GetDrug(d domain.DrugSearch) (drug domain.Drug, err error)
	GetAllDrug()(drugs []domain.Drug,err error)
	GetName(userId int,Error error) (name string, err error)
	SaveMessage(userId string,isAi bool,message string) (id int,err error)
	GetAllMessages(userId int )(messages []domain.Message,err error)
	GetDrugByType(ctx context.Context,tip string)(drugs []domain.DrugWithoutType,err error)
	GetAllTypes(ctx context.Context)(Types []domain.DrugByType,err error)
}
type INewsUseCase interface{
	GetAll(ctx context.Context,query utils.PaginationQuery)(news *domain.NewsList,err error)
	GetOneById(ctx context.Context,id string) (new *domain.NewWithSinglePhoto,err error)
}

type IDoctorUsecase interface{
	GetAll(ctx context.Context) ([]*domain.DoctorByType,error)
	GetOneByID(ctx context.Context,name string)  ([]*domain.DoctorWithType,error)
}

type IScheduleUseCase interface{
	Create(ctx context.Context,schedule *domain.Schedule) error
	GetAll(ctx context.Context,userId int)([]*domain.ScheduleResponse,error)
}


func NewUserUsecase(repo repository.Repo, bot Bot.Bot) Usecase {
	return &usecase{repo: repo, bot: bot}
}
