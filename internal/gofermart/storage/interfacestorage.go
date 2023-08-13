package storage

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Repository interface {
	registerUser(ctx context.Context, login, password string) error
}

func NewStorage(ctx context.Context, migratePath string, postgresDSN string) (Repository, *PostgresDB, error) {

	DB, err := InitDB(postgresDSN, migratePath)
	if err != nil {
		log.Error("Error in initialization db", (err))
		return nil, nil, err
	}
	return DB, DB, nil

}
