package user

import (
	"errors"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) SignUp(user *entities.User) error {
	userDb := FromUseCase(user)

	if err := r.DB.Create(&userDb).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return constant.ErrDuplicatedData
		}
		return err
	}

	user = userDb.ToUseCase()
	return nil
}
