package models

import "github.com/google/uuid"

type Teacher struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	IeltsScore      float64   `json:"ielts_score"`
	Contact         string    `json:"contact"`
	Bio             string    `json:"bio"`
}

type CreateTeacher struct {
	Name            string  `json:"name"`
	ExperienceYears int     `json:"experience_years"`
	IeltsScore      float64 `json:"ielts_score"`
	Contact         string  `json:"contact"`
	Bio             string  `json:"bio"`
}

type Void struct {}