package rest

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/pkg/jwt"
	"DoctorWho/internal/pkg/sms"
	"DoctorWho/internal/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type controller struct {
	usecase usecase.Usecase
	bot     Bot.Bot
}

func NewController(usecase usecase.Usecase) *controller {
	return &controller{usecase: usecase}
}

func (c controller) SignUp(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	var NewUser domain.NewUser
	err := ctx.ShouldBindJSON(&NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError,"Invalid json")
	}
	log.Println(NewUser.PhoneNumber)
	id, err := c.usecase.RegisterUser(&NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError,"Could`nt register")
		
	}
	s.Set("userId", id)
	s.Save()
	ctx.String(http.StatusOK,"verification code sent")
}

func (c controller) Login(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	var User dto.User
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError,"Invalid json")
	}
	exist, id, err := c.usecase.Login(User.Email)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError,"Could`nt login")
	}
	if exist {
		code := sms.GenerateVerificationCode()
		c.bot.SendNotification(code)
		err = sms.SendEmail(User.Email, code)
		c.bot.SendNotification(User.Email)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusUnauthorized,"invalid credentials")
		}
		s.Set("userId", id)
	    s.Save()
	    ctx.String(http.StatusOK,"verification code sent")
		
	} else {
		ctx.String(http.StatusUnauthorized,"invalid credentials")
	}

}
func (c controller) Verification(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	id:= s.Get("userId")
	var message domain.Sms
	err := ctx.ShouldBindJSON(&message)
	if id == nil {
		c.bot.SendErrorNotification(err)
		id = message.UserId
		c.bot.SendNotification(id.(string))
	}
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.Set("Content-Type", "application/json")
		ctx.Status(http.StatusInternalServerError)
		ctx.String(500,"Invalid json")
	}
	match, err := c.usecase.Verify(id.(string), message.Code)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String( http.StatusUnauthorized,"Not matched")
	}
	if match {
		token, err := jwt.CreateToken(id.(string))
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String( http.StatusInternalServerError ,"error occurred: "+err.Error())
			return
		}
		response := map[string]string{
			"access_token": token,
		}
		ctx.JSON(200,response)
	}
}
func (c controller) Logout(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.Set("Content-Type", "application/json")
	ctx.Status(200)
}

func (c controller) FillUserInfo(ctx *gin.Context) {

	
	var UserInfo dto.UserInfo
	err := ctx.ShouldBindJSON(&UserInfo)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(406, gin.H{
			"Message": "Invalid credentials",
		})
	}
	s:=sessions.Default(ctx)
	UserInfo.Id=s.Get("userId").(int)
	id, err := c.usecase.FillInfo(UserInfo)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(400, gin.H{
			"Message": "Bad request",
		})
	}
	ctx.JSON(200, gin.H{
		"Message": "success",
		"Info id": id,
	})
}
func (c controller) UpdateUserInfo(ctx *gin.Context){
	var User dto.UserInfo
	err:=ctx.ShouldBindJSON(&User)
	if err!=nil{
		c.bot.SendErrorNotification(err)
		ctx.JSON(406, gin.H{
			"Message": "Invalid credentials",
		})
	}
	s:=sessions.Default(ctx)
	User.Id=s.Get("userId").(int)
	id, err:=c.usecase.UpdateInfo(User)
	if err!=nil{
		c.bot.SendErrorNotification(err)
		ctx.String(400,"internal error")
	}
	ctx.String(200,"id: ",id)
}
func (c controller) ShowUserInfo(ctx *gin.Context){
	var User dto.UserInfo
	s:=sessions.Default(ctx)
	User.Id=s.Get("userId").(int)
	User ,err:=c.usecase.GetUserInfo(User.Id)
	if err!=nil{
		
			c.bot.SendErrorNotification(err)
			ctx.String(400,"internal error")
		
	}
	ctx.JSON(200,User)
}
func (c controller) GetProgress(ctx *gin.Context) {
	var progress dto.Progress
	s:=sessions.Default(ctx)
	progress.UserId=s.Get("userId").(int)
}

func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	user := c.usecase.GetAll()
	log.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
