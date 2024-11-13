package handlers

import (
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type GalleryHandler struct {
	repo *repository.GalleryRepository
}

func NewGalleryHandler(repo *repository.GalleryRepository) *GalleryHandler {
	return &GalleryHandler{repo: repo}
}
