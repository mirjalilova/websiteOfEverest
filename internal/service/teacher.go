package service

import (
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
