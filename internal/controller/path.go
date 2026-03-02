package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"go-file-share/configs"
	"go-file-share/internal/resource"
	"go-file-share/internal/service"

	"github.com/gin-gonic/gin"
)

type PathController interface {
	ShowFolder(gc *gin.Context)
	DownloadFile(gc *gin.Context)
	UploadFile(gc *gin.Context)
}

type pathController struct {
	dirPath     string
	recursive   bool
	pathService service.PathService
}

func (pc *pathController) ShowFolder(gc *gin.Context) {
	fullPath := filepath.Join(pc.dirPath, filepath.Clean(gc.Request.URL.Path))
	if !pc.recursive && gc.Request.URL.Path != "/" {
		rel, _ := filepath.Rel(pc.dirPath, fullPath)
		if strings.Contains(rel, string(os.PathSeparator)) {
			http.Error(gc.Writer, "Subfolder access disabled (-R not set)", http.StatusForbidden)
			return
		}
	}
	if !strings.HasPrefix(fullPath, pc.dirPath) {
		http.Error(gc.Writer, "Access Denied", http.StatusForbidden)
		return
	}
	pathInfo := pc.pathService.ListPathInfo(fullPath, gc.Request.URL.Path)
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"pathInfo": pathInfo,
		"version":  configs.Version,
	})
}

func (pc *pathController) DownloadFile(gc *gin.Context) {
	fullPath := filepath.Join(pc.dirPath, filepath.Clean(strings.Replace(gc.Request.URL.Path, "/files", "", 1)))
	http.ServeFile(gc.Writer, gc.Request, fullPath)
}

func (pc *pathController) UploadFile(gc *gin.Context) {
	file, handler, err := gc.Request.FormFile("file")
	if err != nil {
		http.Error(gc.Writer, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileUpload := resource.FileUpload{
		UploadDir: filepath.Join(pc.dirPath, filepath.Clean(strings.Replace(gc.Request.URL.Path, "/files", "", 1))),
		File:      file,
		Filename:  handler.Filename,
	}

	err = pc.pathService.UploadFile(fileUpload)
	if err != nil {
		http.Error(gc.Writer, "Error saving the file", http.StatusInternalServerError)
		return
	}

	gc.Redirect(http.StatusSeeOther, filepath.Clean(strings.Replace(gc.Request.URL.Path, "/files", "", 1)))
}

func NewPathController(dirPath string, recursive bool, pathService service.PathService) PathController {
	return &pathController{
		dirPath:     dirPath,
		recursive:   recursive,
		pathService: pathService,
	}

}
