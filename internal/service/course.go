package service

import (
	"context"
	"log/slog"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type CourseService struct {
	storage r.Storage
	pb.UnimplementedCourseServiceServer
}

func NewdCourseService(repo *r.Storage) *CourseService {
	return &CourseService{storage: *repo}
}

func (s *CourseService) Create(ctx context.Context, req *pb.CreateCourse) (*pb.Void, error) {
	res, err := s.storage.CourseS.Create(req)
	if err != nil {
		slog.Error("Error creating course", "err", err)
		return nil, err
	}

	slog.Info("Created course")
	return res, nil
}

func (s *CourseService) Update(ctx context.Context, req *pb.UpdateCourse) (*pb.Void, error) {
	res, err := s.storage.CourseS.Update(req)
    if err != nil {
        slog.Error("Error updating course", "err", err)
        return nil, err
    }

    slog.Info("Updated course")
    return res, nil
}

func (s *CourseService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.CourseS.Delete(req)
    if err != nil {
        slog.Error("Error deleting course", "err", err)
        return nil, err
    }

    slog.Info("Deleted course")
    return res, nil
}

func (s *CourseService) GetById(ctx context.Context, req *pb.ById) (*pb.CourseRes, error) {
	res, err := s.storage.CourseS.GetById(req)
    if err != nil {
        slog.Error("Error getting course by id", "err", err)
        return nil, err
    }

    slog.Info("Course retrieved")
    return res, nil
}

func (s *CourseService) GetList(ctx context.Context, req *pb.GetListCourseReq) (*pb.GetListCourseRes, error) {
	res, err := s.storage.CourseS.GetList(req)
    if err!= nil {
        slog.Error("Error getting courses list", "err", err)
        return nil, err
    }

    slog.Info("Retrieved courses list")
    return res, nil
}
