package postgresqlgorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabasePostgress(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
