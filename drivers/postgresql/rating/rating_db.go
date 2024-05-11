package rating

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	ID         uuid.UUID `gorm:"type:varchar(100);"`
	FromUserID uuid.UUID `gorm:"type:varchar(100);"`
	ToUserID   uuid.UUID `gorm:"type:varchar(100);"`
	JobID      uuid.UUID `gorm:"type:varchar(100);"`
	Star       int       `gorm:"type:int;"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt
}

func FromUseCase(rating *entities.Rating) *Rating {
	return &Rating{
		ID:         rating.ID,
		FromUserID: rating.FromUserID,
		ToUserID:   rating.ToUserID,
		JobID:      rating.JobID,
		Star:       rating.Star,
		CreatedAt:  rating.CreatedAt,
		UpdatedAt:  rating.UpdatedAt,
	}
}

func (r *Rating) ToUseCase() *entities.Rating {
	return &entities.Rating{
		ID:         r.ID,
		FromUserID: r.FromUserID,
		ToUserID:   r.ToUserID,
		JobID:      r.JobID,
		Star:       r.Star,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
