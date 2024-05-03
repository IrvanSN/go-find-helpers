package thumbnail

import "github.com/google/uuid"

type Thumbnail struct {
	ID          uuid.UUID `gorm:"type:varchar(100);"`
	JobID       uuid.UUID `gorm:"type:varchar(100);"`
	ImageKey    string    `gorm:"type:text;"`
	Description string    `gorm:"type:varchar(100);"`
}
