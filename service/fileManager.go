package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type IFileManager interface {
	PutFile(header *multipart.FileHeader) error
	FetchFile()
}

func NewLocalFileManger(path string) *LocalFileManager {
	return &LocalFileManager{
		uploadPath: path,
	}
}

type LocalFileManager struct {
	uploadPath string
}

func (r *LocalFileManager) PutFile(header *multipart.FileHeader) error {
	src, err := header.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(r.getFileName(header.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func (r *LocalFileManager) getFileName(name string) string {
	return fmt.Sprintf("%s/%d_%s", r.uploadPath, time.Now().Unix(), name)
}

func (r *LocalFileManager) FetchFile() {

}


