package upload

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		localTargetFilePath := "./public/" + filename
		if err := c.SaveUploadedFile(file, localTargetFilePath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
