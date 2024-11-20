package service

import (
	"context"
	"log/slog"

	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type BranchesService struct {
	storage r.Storage
	pb.UnimplementedBranchesServiceServer
}

func NewdBranchesService(repo *r.Storage) *BranchesService {
	return &BranchesService{storage: *repo}
}

func (s *BranchesService) Create(ctx context.Context, req *pb.CreateBranches) (*pb.Void, error) {
	res, err := s.storage.BranchesS.Create(req)
	if err != nil {
		slog.Error("Error creating branches", "err", err)
		return nil, err
	}

	slog.Info("Created branches")
	return res, nil
}

func (s *BranchesService) Update(ctx context.Context, req *pb.UpdateBranches) (*pb.Void, error) {
	res, err := s.storage.BranchesS.Update(req)
	if err != nil {
		slog.Error("Error updating branches", "err", err)
		return nil, err
	}

	slog.Info("Updated branches")
	return res, nil
}

func (s *BranchesService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	res, err := s.storage.BranchesS.Delete(req)
	if err != nil {
		slog.Error("Error deleting branches", "err", err)
		return nil, err
	}

	slog.Info("Deleted branches")
	return res, nil
}

func (s *BranchesService) GetById(ctx context.Context, req *pb.ById) (*pb.BranchesRes, error) {
	res, err := s.storage.BranchesS.GetById(req)
	if err != nil {
		slog.Error("Error getting branches by id", "err", err)
		return nil, err
	}

	slog.Info("Retrieved branches by id")
	return res, nil
}

func (s *BranchesService) GetList(ctx context.Context, req *pb.GetListBranchesReq) (*pb.GetListBranchesRes, error) {
	res, err := s.storage.BranchesS.GetList(req)
	if err != nil {
		slog.Error("Error getting branches list", "err", err)
		return nil, err
	}

	slog.Info("Retrieved branches list")
	return res, nil
}
