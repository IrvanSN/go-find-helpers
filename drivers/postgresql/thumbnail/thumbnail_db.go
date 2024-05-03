package thumbnail

import "github.com/google/uuid"

type Thumbnail struct {
	ID          uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	JobID       uuid.UUID `gorm:"type:varchar(100);" json:"job_id"`
	ImageKey    string    `gorm:"type:text;" json:"image_key"`
	Description string    `gorm:"type:varchar(100);" json:"description"`
}
