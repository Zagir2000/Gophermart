package handlers

import (
	"context"
	"net/http"

	"github.com/MlDenis/internal/gofermart/storage"
)

type MetricHandlerDB struct {
	Storage storage.Repository
	pgDB    *storage.PostgresDB
}

func MetricHandlerNew(s storage.Repository, pgDB *storage.PostgresDB) *MetricHandlerDB {
	return &MetricHandlerDB{
		Storage: s,
		pgDB:    pgDB,
	}
}

func (m *MetricHandlerDB) RegisterNewUser(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

	}

}
