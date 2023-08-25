package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MlDenis/internal/accrual/models"
	log "github.com/sirupsen/logrus"
)

// Регистрация информации о вознаграждении за товар
func (m *HandlerDB) RegisterInfoReward(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Error("got request with bad method", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("Content-Type") != "application/json" {
			log.Error("wrong Content-Type", req.Header.Get("Content-Type"))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonGoods := &models.Reward{}
		dec := json.NewDecoder(req.Body)
		if err := dec.Decode(jsonGoods); err != nil {
			log.Error("cannot decode request JSON body", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		err := m.Storage.RegisterInfoInDB(ctx, jsonGoods)
		if err != nil {
			log.Error("error in add info for goods in db", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusAccepted)
		return
	}
}
