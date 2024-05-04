package request

import "github.com/google/uuid"

type JobCreate struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	FromAddressId  uuid.UUID `json:"from_address_id"`
	ToAddressId    uuid.UUID `json:"to_address_id"`
	HelperRequired uint      `json:"helper_required"`
	CategoryId     uuid.UUID `json:"category_id"`
}
