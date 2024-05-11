package rating

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	ID        uuid.UUID `gorm:"type:varchar(100);"`
	UserID    uuid.UUID `gorm:"type:varchar(100);"`
	JobID     uuid.UUID `gorm:"type:varchar(100);"`
	Star      int       `gorm:"type:int;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
