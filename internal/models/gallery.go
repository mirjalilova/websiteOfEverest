package models

import "github.com/google/uuid"

type Gallery struct {
    ID          uuid.UUID `json:"id"`
    PictureURL  string    `json:"picture_url"`
}
