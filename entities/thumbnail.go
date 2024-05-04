package entities

import "github.com/google/uuid"

type Thumbnail struct {
	ID          uuid.UUID
	ImageKey    string
	Description string
}
