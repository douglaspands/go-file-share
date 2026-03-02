package resource

import "mime/multipart"

type FileUpload struct {
	UploadDir string
	File      multipart.File
	Filename  string
}
