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
	ID             uuid.UUID `gorm:"type:varchar(100);" json:"id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	Reward         float64   `gorm:"type:decimal;not null" json:"reward"`
	FromAddressID  uuid.UUID `gorm:"type:varchar(100)" json:"from_address_id"`
	FromAddress    address.Address
	ToAddressID    uuid.UUID `gorm:"type:varchar(100)" json:"to_address_id"`
	ToAddress      address.Address
	Status         string    `gorm:"type:varchar(50);not null" json:"status"`
	HelperRequired uint      `gorm:"not null" json:"helper_required"`
	CategoryID     uuid.UUID `gorm:"type:varchar(100);not null" json:"category_id"`
	Category       category.Category
	UserID         uuid.UUID `gorm:"type:varchar(100);not null" json:"user_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Rewards        []reward.Reward
	Thumbnails     []thumbnail.Thumbnail
}
