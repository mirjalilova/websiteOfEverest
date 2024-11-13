package service

import (
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	r "github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type DashboardService struct {
	storage r.Storage
	pb.UnimplementedDashboardServiceServer
}

func NewDashboardService(repo *r.Storage) *DashboardService {
	return &DashboardService{storage: *repo}
}
