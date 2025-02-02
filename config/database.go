package config

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	config := Config
	enCodedPassword := url.QueryEscape(config.Database.Password)
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		enCodedPassword,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(config.Database.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(config.Database.MaxOpenConnection)
	sqlDb.SetConnMaxLifetime(time.Duration(config.Database.MaxIdleConnection) * time.Second)
	sqlDb.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime))
	return db, nil
}
