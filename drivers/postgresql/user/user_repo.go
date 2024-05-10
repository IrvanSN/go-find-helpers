package user

import (
	"errors"
	"fmt"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	userDb := FromUseCase(user)

	if err := r.DB.Model(&userDb).Association("Addresses").Append(&userDb); err != nil {
		fmt.Println("err", err)
		return constant.ErrInsertDatabase
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) GetAllAddresses(user *entities.User) error {
	userDb := FromUseCase(user)

	if err := r.DB.Preload("Addresses").First(&userDb, "id = ?", userDb.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) Find(user *entities.User) error {
	userDb := FromUseCase(user)

	if err := r.DB.Preload(clause.Associations).First(&userDb, userDb.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) Update(user *entities.User) error {
	userDb := FromUseCase(user)

	db := r.DB.Where("id = ?", userDb.ID).Updates(&userDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(user *entities.User) error {
	userDb := FromUseCase(user)

	db := r.DB.Where("id = ?", userDb.ID).Delete(&userDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*user = *userDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(user *[]entities.User) error {
	var userDb []User

	if err := r.DB.Preload("Auth").Find(&userDb).Error; err != nil {
		return err
	}

	for _, _user := range userDb {
		*user = append(*user, *_user.ToUseCase())
	}
	return nil
}
