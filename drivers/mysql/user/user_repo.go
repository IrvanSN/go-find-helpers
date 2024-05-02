package user

import (
	"github.com/google/uuid"
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
	userDb.ID = uuid.New()

	if err := r.DB.Create(&userDb).Error; err != nil {
		return err
	}

	user = userDb.ToUseCase()
	return nil
}
