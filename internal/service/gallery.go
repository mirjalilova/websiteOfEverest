package service

import (
	"context"
	"log/slog"

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

func (s *GalleryService) Create(ctx context.Context, req *pb.CreateGallery) (*pb.Void, error) {
	res, err := s.storage.GalleryS.Create(req)
	if err != nil {
		slog.Error("Error creating gallery", "err", err)
		return nil, err
	}

	slog.Info("Created gallery")
	return res, nil
}

func (s *GalleryService) Update(ctx context.Context, req *pb.UpdateGallery) (*pb.Void, error) {
	res, err := s.storage.GalleryS.Update(req)
	if err != nil {
		slog.Error("Error updating gallery", "err", err)
		return nil, err
	}

	slog.Info("Updated gallery")
	return res, nil
}

func (s *GalleryService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.GalleryS.Delete(req)
	if err != nil {
		slog.Error("Error deleting gallery", "err", err)
		return nil, err
	}

	slog.Info("Deleted gallery")
	return res, nil
}

func (s *GalleryService) GetById(ctx context.Context, req *pb.ById) (*pb.GalleryRes, error) {
	res, err := s.storage.GalleryS.GetById(req)
	if err != nil {
		slog.Error("Error getting gallery by id", "err", err)
		return nil, err
	}

	slog.Info("Retrieved gallery by id")
	return res, nil
}

func (s *GalleryService) GetList(ctx context.Context, req *pb.GetListGalleryReq) (*pb.GetListGalleryRes, error) {
	res, err := s.storage.GalleryS.GetList(req)
	if err != nil {
		slog.Error("Error getting gallery list", "err", err)
		return nil, err
	}

	slog.Info("Retrieved gallery list")
	return res, nil
}
