package storage

import (
	"fmt"
	"image"
)

type Storage interface {
	SetImage(img image.Image) error
	GetBaseImage() (error, image.Image)
	GetWorkingImage() (error, image.Image)
}

type storage struct {
	img          image.Image
	workingImage image.Image
}

func (s *storage) SetImage(img image.Image) error {
	s.img = img
	s.workingImage = img

	return nil
}

func (s *storage) GetBaseImage() (error, image.Image) {
	if s.img == nil || s.workingImage == nil {
		return fmt.Errorf("Image not set. Call SetImage method"), nil
	}

	return nil, s.img
}

func (s *storage) GetWorkingImage() (error, image.Image) {
	if s.img == nil || s.workingImage == nil {
		return fmt.Errorf("Image not set. Call SetImage method"), nil
	}

	return nil, s.workingImage
}

func New() *storage {
	return &storage{}
}
