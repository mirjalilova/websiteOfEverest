package service

import (
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type GalleryService struct {
	storage r.Storage
	pb.UnimplementedGalleryServiceServer
}

func NewGalleryService(repo *r.Storage) *GalleryService {
	return &GalleryService{storage: *repo}
}
