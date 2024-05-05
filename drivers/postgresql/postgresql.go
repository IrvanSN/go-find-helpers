package postgresql

import (
	"fmt"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/auth"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/job"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/payment"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/rating"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/thumbnail"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		config.DBHost,
		config.DBUser,
		config.DBPass,
		config.DBName,
		config.DBPort,
	)
	fmt.Println("Connecting to", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	MigrationUser(db)
	return db
}

func MigrationUser(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&auth.Auth{},
		&address.Address{},
		&category.Category{},
		&job.Job{},
		&rating.Rating{},
		&transaction.Transaction{},
		&payment.Payment{},
		&thumbnail.Thumbnail{},
	)
	if err != nil {
		return
	}
}
