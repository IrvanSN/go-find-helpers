package job

import (
	"github.com/irvansn/go-find-helpers/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewJobRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(job *entities.Job) error {
	return nil
}
