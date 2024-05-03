package rating

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	ID        uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	UserID    uuid.UUID `gorm:"type:varchar(100);" json:"user_id"`
	Star      int       `gorm:"type:int;" json:"star"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
