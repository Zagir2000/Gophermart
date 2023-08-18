package handlers

import (
	"github.com/MlDenis/internal/gofermart/auth/cache"
	"github.com/MlDenis/internal/gofermart/storage"
)

// структура для наших хэндлеров, далее надо будет добавить возмонжо логер и тд
type HandlerDB struct {
	Storage storage.DBInterface
	DataJWT *cache.DataJWT
}

func HandlerNew(s storage.DBInterface, DataJWT *cache.DataJWT) *HandlerDB {
	return &HandlerDB{
		Storage: s,
		DataJWT: DataJWT,
	}
}
