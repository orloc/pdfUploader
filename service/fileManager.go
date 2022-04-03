package service

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image/jpeg"
	"math/rand"
	"mime/multipart"
	"os"
)

type IFileManager interface {
	PutFile(header *multipart.FileHeader) ([]string, error)
}

func NewLocalFileManger(path string) *LocalFileManager {
	return &LocalFileManager{
		uploadPath: path,
	}
}

type LocalFileManager struct {
	uploadPath string
}

func (r *LocalFileManager) PutFile(header *multipart.FileHeader) ([]string, error) {
	src, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	imgs, err := r.generateImagesFromPDF(src)
	if err != nil {
		return nil, err
	}
	return imgs, nil
}

func (r *LocalFileManager) generateImagesFromPDF(src multipart.File) ([]string, error) {
	var locations []string

	doc, err := fitz.NewFromReader(src)
	defer doc.Close()
	if err != nil {
		return nil, err
	}

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return nil, err
		}
		fN := r.getFileName(n)
		f, err := os.Create(fN)

		if err != nil {
			return nil, err
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return nil, err
		}
		locations = append(locations, fN)
		f.Close()
	}

	return locations, nil
}

func (r *LocalFileManager) getFileName(n int) string {
	return fmt.Sprintf("%s/%d_%s.jpg", r.uploadPath, n, r.randomSuffString(5))
}

func (r *LocalFileManager) randomSuffString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return  string(s)
}

