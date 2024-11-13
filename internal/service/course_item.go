package service

import (
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
