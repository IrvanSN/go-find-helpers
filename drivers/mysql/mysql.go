package mysql

import (
	"fmt"
	"github.com/irvansn/go-find-helpers/drivers/mysql/address"
	"github.com/irvansn/go-find-helpers/drivers/mysql/auth"
	"github.com/irvansn/go-find-helpers/drivers/mysql/user"
	"gorm.io/driver/mysql"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	MigrationUser(db)
	return db
}

func MigrationUser(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{}, &auth.Auth{}, address.Address{})
	if err != nil {
		return
	}
}
