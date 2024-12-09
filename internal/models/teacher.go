package models

import "github.com/google/uuid"

type Teacher struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears string       `json:"experience_years"`
	IeltsScore      float64   `json:"ielts_score"`
	Contact         string    `json:"contact"`
}

type CreateTeacher struct {
	Name            string  `json:"name"`
	ExperienceYears string     `json:"experience_years"`
	IeltsScore      float64 `json:"ielts_score"`
	Contact         string  `json:"contact"`
	GraduatedStudents string `json:"graduated_students"`}

type Void struct {}