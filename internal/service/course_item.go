package service

import (
	"context"
	"log/slog"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type CourseItemService struct {
	storage r.Storage
	pb.UnimplementedCourseItemServiceServer
}

func NewCourseItemService(repo *r.Storage) *CourseItemService {
	return &CourseItemService{storage: *repo}
}

func (s *CourseItemService) Create(ctx context.Context, req *pb.CreateCourseItem) (*pb.Void, error) {
	res, err := s.storage.CourseItemS.Create(req)
	if err != nil {
		slog.Error("Error creating course item", "err", err)
		return nil, err
	}

	slog.Info("Created course item")
	return res, nil
}

func (s *CourseItemService) Update(ctx context.Context, req *pb.UpdateCourseItem) (*pb.Void, error) {
	res, err := s.storage.CourseItemS.Update(req)
	if err != nil {
		slog.Error("Error updating course item", "err", err)
		return nil, err
	}

	slog.Info("Updated course item")
	return res, nil
}

func (s *CourseItemService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.CourseItemS.Delete(req)
	if err != nil {
		slog.Error("Error deleting course item", "err", err)
		return nil, err
	}

	slog.Info("Deleted course item")
	return res, nil
}

func (s *CourseItemService) GetById(ctx context.Context, req *pb.ById) (*pb.CourseItemRes, error) {
	res, err := s.storage.CourseItemS.GetById(req)
	if err != nil {
		slog.Error("Error getting course item by id", "err", err)
		return nil, err
	}

	slog.Info("Retrieved course item by id")
	return res, nil
}

func (s *CourseItemService) GetList(ctx context.Context, req *pb.GetListCourseItemReq) (*pb.GetListCourseItemRes, error) {
	res, err := s.storage.CourseItemS.GetList(req)
	if err != nil {
		slog.Error("Error getting course items list", "err", err)
		return nil, err
	}

	slog.Info("Retrieved course items list")
	return res, nil
}
