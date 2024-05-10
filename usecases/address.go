package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
)

type AddressUseCase struct {
	repository entities.AddressRepositoryInterface
}

func NewAddressUseCase(repository entities.AddressRepositoryInterface) *AddressUseCase {
	return &AddressUseCase{repository: repository}
}

func (c *AddressUseCase) Update(address *entities.Address, user *middlewares.Claims) (entities.Address, error) {
	if address.ID == uuid.Nil {
		return entities.Address{}, constant.ErrEmptyInput
	}

	if user.Role != "ADMIN" {
		return entities.Address{}, constant.ErrNotAuthorized
	}

	if err := c.repository.Update(address); err != nil {
		return entities.Address{}, err
	}

	if err := c.repository.Get(address); err != nil {
		return entities.Address{}, err
	}

	return *address, nil
}

func (c *AddressUseCase) Delete(address *entities.Address, user *middlewares.Claims) (entities.Address, error) {
	if address.ID == uuid.Nil {
		return entities.Address{}, constant.ErrEmptyInput
	}

	if user.Role != "ADMIN" {
		return entities.Address{}, constant.ErrNotAuthorized
	}

	if err := c.repository.Delete(address); err != nil {
		return entities.Address{}, err
	}

	return *address, nil
}
