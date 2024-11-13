package models

import "github.com/google/uuid"

type Course struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Duration float64   `json:"duration"`
}
