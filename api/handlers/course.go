package handlers

import (
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type CourseHandler struct {
	repo *repository.CourseRepository
}

func NewCourseHandler(repo *repository.CourseRepository) *CourseHandler {
	return &CourseHandler{repo: repo}
}
