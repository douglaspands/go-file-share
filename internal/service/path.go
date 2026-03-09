package service

import (
	resource "go-file-share/internal/resource"
	utils "go-file-share/internal/utils"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type PathService interface {
	ListPathInfo(fullPath string, urlPath string, recursive bool) *resource.PathInfoList
	UploadFile(fileUpload resource.FileUpload) error
}

type pathService struct {
}

func (ps *pathService) ListPathInfo(fullPath string, urlPath string, recursive bool) *resource.PathInfoList {
	entries, _ := os.ReadDir(fullPath)
	var paths []resource.PathInfo
	if urlPath != "/" {
		parent := filepath.Dir(strings.TrimSuffix(urlPath, "/"))
		paths = append(paths, resource.PathInfo{Name: "../", IsDir: true, Path: filepath.ToSlash(parent)})
	}
	for _, entry := range entries {
		pathName := entry.Name()
		isDir := entry.IsDir()
		if !recursive && isDir {
			continue
		}
		dirPath := urlPath
		displaySize := ""
		if !isDir {
			info, _ := entry.Info()
			displaySize = utils.FormatBytes(info.Size())
			dirPath = filepath.Join("/files", dirPath)
		}
		if !strings.HasPrefix(pathName, ".") {
			path := filepath.Join(dirPath, pathName)
			paths = append(paths, resource.PathInfo{
				Name:  pathName,
				IsDir: isDir,
				Size:  displaySize,
				Path:  filepath.ToSlash(path),
			})
		}
	}
	sort.Slice(paths, func(i, j int) bool {
		if paths[i].IsDir != paths[j].IsDir {
			return paths[i].IsDir
		}
		return strings.ToLower(paths[i].Name) < strings.ToLower(paths[j].Name)
	})
	return &resource.PathInfoList{
		CurrentPath: urlPath,
		Paths:       paths,
	}
}

func (ps pathService) UploadFile(fileUpload resource.FileUpload) error {
	defer fileUpload.File.Close()
	dstPath := filepath.Join(fileUpload.UploadDir, fileUpload.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, fileUpload.File); err != nil {
		return err
	}
	return nil
}

func NewPathService() PathService {
	return &pathService{}
}
