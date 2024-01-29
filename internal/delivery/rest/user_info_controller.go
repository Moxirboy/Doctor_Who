package rest

import (
	"net/http"
	"testDeployment/internal/delivery/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (c controller) FillUserInfo(ctx *gin.Context) {
	var UserInfo dto.UserInfo
	err := ctx.ShouldBindJSON(&UserInfo)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(406, gin.H{
			"Message": "Invalid credentials",
		})
		return
	}
	s := sessions.Default(ctx)
	UserInfo.Id = s.Get("userId").(int)
	id, err := c.usecase.FillInfo(UserInfo)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(400, gin.H{
			"Message": "Bad request",
		})
		return
	}
	s.Set("name", UserInfo.Name)
	s.Set("age", UserInfo.Age)
	s.Set("gender", UserInfo.Gender)
	s.Save()
	ctx.JSON(200, gin.H{
		"Message": "success",
		"Info id": id,
	})
}
func (c controller) UpdateUserInfo(ctx *gin.Context) {
	var User dto.UserInfo
	s := sessions.Default(ctx)
	User.Id = s.Get("userId").(int)
	err := ctx.ShouldBindJSON(&User)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(406, gin.H{
			"Message": "Invalid credentials",
		})
		return
	}

	id, err := c.usecase.UpdateInfo(User)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.String(400, "internal error")
		return
	}
	ctx.String(200, "id: ", id)
}
func (c controller) ShowUserInfo(ctx *gin.Context) {
	var User dto.UserInfo
	s := sessions.Default(ctx)
	User.Id = s.Get("userId").(int)
	if User.Id == 0 {
		ctx.String(http.StatusUnauthorized, "Not registered")
		return
	}
	User, err := c.usecase.GetUserInfo(User.Id)
	if err != nil {
		c.bot.SendErrorNotification(err)
		ctx.JSON(400, gin.H{
			"message": err})
		return

	}
	ctx.JSON(200, User)
}
