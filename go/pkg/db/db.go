package db

import (
	"github.com/jinzhu/gorm"
	// Register postgres driver
	_ "github.com/lib/pq"
)

const (
	// MigrationsTable is the name of the table tracking applied migrations.
	MigrationsTable = "schema_migrations"
)

type db struct {
	*gorm.DB
}

// New returns a Postgres backed Repo.
func New(dsn string) (*db, error) {
	gormDB, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &db{
		DB: gormDB,
	}, nil
}
