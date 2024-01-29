package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karrick/godirwalk"
)
func (c controller) GetPhoto(ctx *gin.Context){

	filePath := ctx.Query("path")

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	_, fileName := filepath.Split(filePath)

	// Set the appropriate headers for displaying file content in the browser
	ctx.Header("Content-Description", "File Content")
	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", fileName))
	ctx.Header("Content-Type", http.DetectContentType(fileContent))
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(fileContent)))

	// Write the file content to the response body
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(fileContent)
}
func (c controller) Download(ctx *gin.Context){
	// Get the file path from the query parameter
	filePath := ctx.Query("path")

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	_, fileName := filepath.Split(filePath)

	// Set the appropriate headers for file download
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Cache-Control", "no-cache")

	// Serve the file
	ctx.File(filePath)
}
func(s controller) Getdirectory(c *gin.Context) {
	dirname := "."
	var tree string
	
    _ = godirwalk.Walk(dirname, &godirwalk.Options{
        Callback: func(osPathname string, de *godirwalk.Dirent) error {
            // Following string operation is not most performant way
            // of doing this, but common enough to warrant a simple
            // example here:
            if strings.Contains(osPathname, ".git") {
                return godirwalk.SkipThis
            }
			tree+=""+osPathname
            fmt.Printf("%s %s\n", de.ModeType(), osPathname)
            return nil
        },
        Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
    })
	c.String(http.StatusOK, tree)
}
