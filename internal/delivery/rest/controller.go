package rest

import (
	request "testDeployment/internal/delivery/http"
	"testDeployment/internal/delivery/middleware"
	"testDeployment/internal/usecase"
	"testDeployment/pkg/Bot"

	"github.com/gin-gonic/gin"
)

type controller struct {
	usecase usecase.Usecase
	bot     Bot.Bot
	http request.CustomJSONRequester
}

func NewController(g *gin.RouterGroup,usecase usecase.Usecase, bot Bot.Bot,request request.CustomJSONRequester )  {
	controller:=controller{
		usecase: usecase,
		bot: bot,
		http: request,
	}
	r:=g.Group("/")
	r.GET("/download",controller.Download)
	r.GET("/path",controller.GetPhoto)
	r.GET("/get" ,controller.Getdirectory)
	save := r.Group("/save")
	{
		save.GET("/", func(c *gin.Context) {
			c.String(200, "Hello from save")
		})
		
		drugs := save.Group("/drugs")
		{
			drugs.GET("/", controller.DrugIndexHandler)
			drugs.POST("/upload", controller.DrugUploadHandler)
		}
	}
	v1Group := r.Group("/inner")
	{
		v1Group.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
		v1Group.POST("/signup", controller.SignUp)
		v1Group.POST("/login", controller.Login)

		v1Group.POST("/verification", controller.Verification)
		dash := v1Group.Group("/dashboard")
		{
			dash.GET("/", func(c *gin.Context) {
				c.String(200, "Hello from dashboard")
			})
		
			dash.GET("/searchDrug", controller.SearchDrug)
			dash.GET("/getOneDrugById", controller.GetDrug)
			dash.GET("/getdrug",controller.GetAllDrug)
			dash.GET("/getbytype",controller.GetDrugByType)
			dash.GET("/getalltype",controller.GetAllTypes)

			middle := dash.Group("/middle")
			middle.Use(middleware.AuthMiddleware())
			{
				middle.POST("/send-request",controller.SendMessage)
				middle.GET("/get-all-messages",controller.GetAllMessages)
				middle.POST("/updateuserinfo", controller.UpdateUserInfo)
				middle.GET("/showUserInfo", controller.ShowUserInfo)
				middle.GET("/logout", controller.Logout)
				middle.GET("/deleteAccount", controller.DeleteAccount)
			}
			{
			}
			dash.POST("/fillUserInfo", controller.FillUserInfo)
		}
}}