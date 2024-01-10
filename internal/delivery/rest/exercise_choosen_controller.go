package rest

import (

	"DoctorWho/internal/domain"

	"github.com/gin-gonic/gin"
)

func (c controller) MarkAsDone(ctx *gin.Context) {

	var Mark domain.MarkAsDone
	err := ctx.ShouldBindJSON(&Mark)
	if err != nil {
		ctx.JSON(406, gin.H{
			"Message": "Invalid credentials",
		})
		return
	}
	id, err := c.usecase.MarkAsDone(Mark)
	if err != nil {
		ctx.JSON(406, gin.H{
			"Message": "try again",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"id ": id,
	})
}
