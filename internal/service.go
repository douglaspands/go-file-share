package internal

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

func renderDirectory(w http.ResponseWriter, fullPath string, urlPath string) {
	entries, _ := os.ReadDir(fullPath)

	var files []FileInfo

	if urlPath != "/" {
		parent := filepath.Dir(strings.TrimSuffix(urlPath, "/"))
		files = append(files, FileInfo{Name: "../", IsDir: true, Path: parent})
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		displaySize := ""
		if !entry.IsDir() {
			displaySize = formatBytes(info.Size())
		}
		files = append(files, FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Size:  displaySize,
			Path:  filepath.Join(urlPath, entry.Name()),
		})
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir != files[j].IsDir {
			return files[i].IsDir
		}
		return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
	})

	t := template.Must(template.New("listing").Parse(htmlTemplate))
	t.Execute(w, struct {
		CurrentPath string
		Files       []FileInfo
	}{
		CurrentPath: urlPath,
		Files:       files,
	})
}

func handleUpload(w http.ResponseWriter, r *http.Request, uploadDir string) {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	dstPath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
}
