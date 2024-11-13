package service

import (
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
