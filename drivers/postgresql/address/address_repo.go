package address

import (
	"errors"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Update(address *entities.Address) error {
	addressDb := FromUseCase(address)

	db := r.DB.Where("id = ?", addressDb.ID).Updates(&addressDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*address = *addressDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(address *entities.Address) error {
	addressDb := FromUseCase(address)

	db := r.DB.Delete(&addressDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*address = *addressDb.ToUseCase()
	return nil
}

func (r *Repo) Get(address *entities.Address) error {
	addressDb := FromUseCase(address)

	if err := r.DB.Find(&addressDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*address = *addressDb.ToUseCase()
	return nil
}
