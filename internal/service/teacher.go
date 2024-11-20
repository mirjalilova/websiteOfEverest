package service

import (
	"context"
	"log/slog"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type TeacherService struct {
	storage r.Storage
	pb.UnimplementedTeacherServiceServer
}

func NewTeacherService(repo *r.Storage) *TeacherService {
	return &TeacherService{storage: *repo}
}

func (s *TeacherService) Create(ctx context.Context, req *pb.CreateTeacher) (*pb.Void, error) {
	res, err := s.storage.TeacherS.Create(req)
	if err != nil {
		slog.Error("Error creating teacher", "err", err)
		return nil, err
	}

	slog.Info("Created teacher")
	return res, nil
}

func (s *TeacherService) Update(ctx context.Context, req *pb.UpdateTeacher) (*pb.Void, error) {
	res, err := s.storage.TeacherS.Update(req)
	if err != nil {
		slog.Error("Error updating teacher", "err", err)
		return nil, err
	}

	slog.Info("Updated teacher")
	return res, nil
}

func (s *TeacherService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.TeacherS.Delete(req)
	if err != nil {
		slog.Error("Error deleting teacher", "err", err)
		return nil, err
	}

	slog.Info("Deleted teacher")
	return res, nil
}

func (s *TeacherService) GetById(ctx context.Context, req *pb.ById) (*pb.TeacherRes, error) {
	res, err := s.storage.TeacherS.GetById(req)
	if err != nil {
		slog.Error("Error getting teacher by id", "err", err)
		return nil, err
	}

	slog.Info("Teacher retrieved")
	return res, nil
}

func (s *TeacherService) GetList(ctx context.Context, req *pb.GetListTeacherReq) (*pb.GetListTeacherRes, error) {
	res, err := s.storage.TeacherS.GetList(req)
	if err != nil {
		slog.Error("Error getting teachers list", "err", err)
		return nil, err
	}

	slog.Info("Retrieved teachers list")
	return res, nil
}
