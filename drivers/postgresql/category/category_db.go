package category

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	Name string    `gorm:"type:varchar(100);not null" json:"name"`
}
