package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

func InitDB(configDB string, migratePath string) (*PostgresDB, error) {
	err := runMigrations(configDB, migratePath)
	if err != nil {
		return nil, fmt.Errorf("failed to run DB migrations: %w", err)
	}
	pool, err := pgxpool.New(context.Background(), configDB)
	if err == nil {
		return &PostgresDB{pool: pool}, nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgerrcode.IsConnectionException(pgErr.Code) {
		log.Error("Database initialization error", err)
		pool, err := pgxpool.New(context.Background(), configDB)
		if err == nil {
			log.Info("Successful database connection")
			return &PostgresDB{pool: pool}, nil
		}

	}
	return nil, fmt.Errorf("failed to create a connection pool: %w", err)
}

func (pgdb *PostgresDB) Close() {
	pgdb.pool.Close()
}

func runMigrations(dsn string, migratePath string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s", "migrations"), dsn)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (dB *PostgresDB) registerUser(ctx context.Context, login, password string) error {

	return nil
}
