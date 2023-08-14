package handlers

import (
	"github.com/MlDenis/internal/gofermart/auth/cache"
	"github.com/MlDenis/internal/gofermart/storage"
)

// структура для наших хэндлеров, далее надо будет добавить возмонжо логер и тд
type HandlerDB struct {
	Storage storage.DBInterface
	pgDB    *storage.PostgresDB
	DataJWT *cache.DataJWT
}

func HandlerNew(s storage.DBInterface, pgDB *storage.PostgresDB, DataJWT *cache.DataJWT) *HandlerDB {
	return &HandlerDB{
		Storage: s,
		pgDB:    pgDB,
		DataJWT: DataJWT,
	}
}
