package request

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
)

type JobUpdateRequest struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	RewardEarned  float64   `json:"reward_earned"`
	FromAddressID uuid.UUID `json:"from_address_id"`
	ToAddressID   uuid.UUID `json:"to_address_id"`
	CategoryID    uuid.UUID `json:"category_id"`
}

func (r *JobUpdateRequest) UpdateFromUseCase() *entities.Job {
	return &entities.Job{
		ID:           r.ID,
		Title:        r.Title,
		Description:  r.Description,
		RewardEarned: r.RewardEarned,
		FromAddress:  entities.Address{ID: r.FromAddressID},
		ToAddress:    entities.Address{ID: r.ToAddressID},
		Category:     entities.Category{ID: r.CategoryID},
	}
}
