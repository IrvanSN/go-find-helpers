package user

import (
	"errors"
	"fmt"
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

func (r *Repo) SignIn(user *entities.User) error {
	userDb := FromUseCase(user)

	if err := r.DB.Joins("Auth").First(&userDb, "\"Auth\".email = ? AND users.role = ?", userDb.Auth.Email, userDb.Role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) AddAddress(user *entities.User) error {
	userDb := AddressFromUseCase(user)
	
	if err := r.DB.Model(&userDb).Association("Addresses").Append(&userDb); err != nil {
		fmt.Println("err", err)
		return constant.ErrInsertDatabase
	}

	*user = *userDb.AddressToUseCase()
	return nil
}
