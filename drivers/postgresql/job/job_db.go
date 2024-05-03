package job

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/reward"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/thumbnail"
	"time"
)

type Job struct {
	ID             uuid.UUID `gorm:"type:varchar(100);"`
	Title          string    `gorm:"type:varchar(255);not null"`
	Description    string    `gorm:"type:text;not null"`
	Reward         float64   `gorm:"type:decimal;not null"`
	FromAddressID  uuid.UUID `gorm:"type:varchar(100)"`
	FromAddress    address.Address
	ToAddressID    uuid.UUID `gorm:"type:varchar(100)"`
	ToAddress      address.Address
	Status         string    `gorm:"type:varchar(50);not null"`
	HelperRequired uint      `gorm:"not null"`
	CategoryID     uuid.UUID `gorm:"type:varchar(100);not null"`
	Category       category.Category
	UserID         uuid.UUID `gorm:"type:varchar(100);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Rewards        []reward.Reward
	Thumbnails     []thumbnail.Thumbnail
}
