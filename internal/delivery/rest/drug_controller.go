package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"testDeployment/internal/delivery/html"
	"testDeployment/internal/domain"
)

func (cr controller) DrugIndexHandler(c *gin.Context) {
	tmpl, err := template.New("index").Parse(html.DrugIndexHTML)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	tmpl.Execute(c.Writer, nil)
}

func (cr controller) DrugUploadHandler(c *gin.Context) {
	var Form domain.Drug
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "Error parsing form")
		return
	}
	Form.Name = c.PostForm("nameOfDrug")
	if Form.Name == "" {
		c.String(http.StatusBadRequest, "name cannot be empty")
		return

	}

	Form.Manufacturer = c.PostForm("manufacturer")
	if Form.Manufacturer == "" {
		c.String(http.StatusBadRequest, "manufacturer cannot be empty")
		return
	}

	Form.Description = c.PostForm("description")
	if Form.Description == "" {
		c.String(http.StatusBadRequest, "description cannot be empty")
		return
	}

	Form.Receipt = c.PostForm("reciept")
	if Form.Receipt == "" {
		c.String(http.StatusBadRequest, "receipt cannot be empty")
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "No files uploaded")
		return
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
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
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer outFile.Close()
		io.Copy(outFile, src)
		Form.Photo = append(Form.Photo, dirPath+"/"+file.Filename)

	}
	_, err = cr.usecase.CreateDrug(Form)
	if err != nil {
		cr.bot.SendErrorNotification(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Redirect to the homepage with the list of uploaded images as query parameters
	c.Status(200)
	c.Redirect(http.StatusSeeOther,"/save/drugs")}
func (c controller) SearchDrug(ctx *gin.Context) {
	var drug domain.DrugSearch
drug.Name = ctx.Query("name")

result, err := c.usecase.GetDrugs(drug)
if err != nil {
    ctx.JSON(406, gin.H{
        "Message": "No such drug",
    })
    return
}

ctx.JSON(200, result)
}
func (c controller) GetDrug(ctx *gin.Context) {
	
	var Drug domain.DrugSearch
	err := ctx.ShouldBindJSON(&Drug)
	Drug.Id=ctx.Query("id")
	drug, err := c.usecase.GetDrug(Drug)
	if err != nil {
		ctx.JSON(406, gin.H{
			"Message": "No such drug",
		})
		return
	}
	ctx.JSON(200, drug)
}
func( c controller) GetAllDrug(ctx *gin.Context){
	drugs,err:=c.usecase.GetAllDrug()
	if err != nil {
		ctx.JSON(406, gin.H{
			"Message": "No  drug",
		})
		return
	}
	ctx.JSON(200, drugs)
}