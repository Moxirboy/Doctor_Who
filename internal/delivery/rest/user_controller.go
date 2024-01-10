package rest

import (
	"net/http"
	"testDeployment/internal/domain"
	"testDeployment/internal/pkg/jwt"
	"testDeployment/internal/pkg/sms"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (c controller) SignUp(ctx *gin.Context) {
	s := sessions.Default(ctx)
	var NewUser domain.NewUser
	err := ctx.ShouldBindJSON(&NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError, "Invalid json")
		return
	}
	if NewUser.PhoneNumber == "" {
		ctx.String(500, "email cannot be empty")
		return
	}
	if NewUser.Password == "" {
		ctx.String(500, "password cannot be empty")
		return
	}
	exist, err := c.usecase.Exist(NewUser)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(409, gin.H{
			"message": "already registered",
		})
		return
	}

	if exist {

		ctx.String(http.StatusNotAcceptable, "email registered")
	} else {
		code := sms.GenerateVerificationCode()

		err = sms.SendEmail(NewUser.PhoneNumber, code)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "Could`nt send code")
			return
		}
		id, err := c.usecase.RegisterUser(&NewUser)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "Could`nt register")
			return
		}
		token, err := jwt.CreateToken(id)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "error occurred: "+err.Error())
			return
		}
		response := map[string]string{
			"access_token": token,
		}

		s.Set("Token", token)
		s.Set("userId", id)
		s.Save()
		ctx.JSON(200, gin.H{
			"message": response,
		})
		
		// c.bot.SendNotification("email:" + NewUser.PhoneNumber)
		// c.bot.SendNotification("code:" + code)
		// s.Set("code", code)
		
		// ctx.String(http.StatusOK, "verification code sent")
	}

}

func (c controller) Login(ctx *gin.Context) {
	s := sessions.Default(ctx)
	var User domain.NewUser
	err := ctx.ShouldBindJSON(&User)

	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError, "Invalid json")
		return
	}
	if User.PhoneNumber == "" {
		ctx.String(500, "email cannot be empty")
		return
	}
	if User.Password == "" {
		ctx.String(500, "password cannot be empty")
		return
	}
	exist, err := c.usecase.Exist(User)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(404, gin.H{
			"message": "No such user",
		})
		return
	}
	
	if !exist {
		ctx.JSON(404, gin.H{
			"message": "No such user",
		})
		ctx.Abort()
		return
	}
	match, id, err := c.usecase.Login(User)

	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(http.StatusInternalServerError, "could not login ")
		return
	}
	if match {
		token, err := jwt.CreateToken(id)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "error occurred: "+err.Error())
			return
		}
		s.Set("Token", token)
		s.Set("userId", id)
		s.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": token,
		})

	} else {
		ctx.String(http.StatusUnauthorized, "Incorrect password")
		return
	}

}
func (c controller) Verification(ctx *gin.Context) {
	s := sessions.Default(ctx)
	id := s.Get("userId")
	var message domain.Sms
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(500, "invalid credentials")
		return
	}
	if message.Code == "" {
		ctx.String(500, "code cannot be empty")
		return
	}
	code := s.Get("code")
	match := code == message.Code
	if match {
		err := c.usecase.UpdateIsVerified(id)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "error occurred: "+err.Error())
			return
		}
		token, err := jwt.CreateToken(id)
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.String(http.StatusInternalServerError, "error occurred: "+err.Error())
			return
		}
		response := map[string]string{
			"access_token": token,
		}

		s.Delete("code")
		s.Set("Token", token)
		s.Save()
		ctx.JSON(200, gin.H{
			"message": response,
		})
	} else {
		ctx.JSON(403, gin.H{
			"message": "Not matched",
		})
	}
}
func (c controller) Logout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.Set("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"message": "successfully logged out",
	})

}

func (cr controller) DeleteAccount(c *gin.Context) {
	s := sessions.Default(c)
	id := s.Get("userId").(int)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not registered"})
		return
	}
	err := cr.usecase.DeleteUser(id)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	s.Delete("userId")
	s.Clear()
	s.Save()
	c.Redirect(http.StatusAccepted, "v1/dashboard")

}
