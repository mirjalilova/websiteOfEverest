package service

import (
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
