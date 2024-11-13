package handlers

import (
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
)

type TeacherHandler struct {
	repo *repository.TeacherRepository
}

func NewTeacherHandler(repo *repository.TeacherRepository) *TeacherHandler {
	return &TeacherHandler{repo: repo}
}
