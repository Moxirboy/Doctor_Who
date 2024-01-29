package rest

import (
	"fmt"
	"io"

	"net/http"
	"os"
	"testDeployment/internal/delivery/dto"
	"testDeployment/pkg/Bot"

	"github.com/gin-gonic/gin"
)


type file struct{
	bot Bot.Bot
}
func NewFileController(g *gin.RouterGroup,bot Bot.Bot){
	controller:=file{bot: bot}
	r:=g.Group("/news")
    r.POST("/upload",controller.Upload)
}
func (f *file) Upload(c *gin.Context){
	var resPhoto dto.PhotoResponse
	form,err:=c.MultipartForm()

	if err != nil {
		f.bot.SendErrorNotification(err)
		c.String(http.StatusBadRequest, "No files uploaded")
		return
	}
	files := form.File["file"]
	

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			f.bot.SendErrorNotification(err)
			c.String(http.StatusInternalServerError, "Error opening file")
			return
		}
		defer src.Close()
	

		// You can save the file to a desired location here
		// For simplicity, we'll just discard the file in this example
		// Uncomment the following lines to save the file:
		dirPath := "storage"
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// Directory doesn't exist, so create it
			err := os.Mkdir(dirPath, 0777)

			if err != nil {
				f.bot.SendErrorNotification(err)
				fmt.Println("Error creating directory:", err)
				return
			}

			fmt.Println("Directory created successfully:", dirPath)
		} else {

			// Directory already exists
			fmt.Println("Directory already exists:", dirPath)
		}

		outFile, err := os.Create(dirPath + "/" + file.Filename)
		if err != nil {
			f.bot.SendErrorNotification(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer outFile.Close()
		io.Copy(outFile, src)
		resPhoto.Path = append(resPhoto.Path, "https://open-data.up.railway.app/api/v1/path?path="+dirPath+"/"+file.Filename)

	}
	
	// Redirect to the homepage with the list of uploaded images as query parameters
	c.Status(200)
	c.JSON(200,resPhoto)
}