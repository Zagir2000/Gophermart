package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MlDenis/internal/gofermart/auth"
	"github.com/MlDenis/internal/gofermart/models"
	log "github.com/sirupsen/logrus"
)

// регистрируем нового пользователя
func (m *HandlerDB) RegisterNewUser(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Debug("got request with bad method", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("Content-Type") != "application/json" {
			log.Debug("wrong Content-Type", req.Header.Get("Content-Type"))
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// десериализуем запрос в структуру модели
		log.Debug("decoding request")
		var jsonUsers models.UserData
		dec := json.NewDecoder(req.Body)
		if err := dec.Decode(&jsonUsers); err != nil {
			log.Debug("cannot decode request JSON body", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonUsers.PasswordHash = auth.HashPassword(jsonUsers.Password)
		err := m.Storage.RegisterUser(ctx, jsonUsers)
		if err != nil {
			err := m.Storage.GetUser(ctx, &jsonUsers)
			if err == nil {
				log.Debug("login busy")
				res.WriteHeader(http.StatusConflict)
			}
			log.Debug("cannot add new counter value")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		err = auth.CreateJwtToken(&jsonUsers)
		if err != nil {
			log.Debug("token not created")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		m.DataJWT.AddToken(&jsonUsers)
		res.Header().Add(models.HeaderHTTP, jsonUsers.Token)
		res.WriteHeader(http.StatusOK)

	}

}

// авторизация
func (m *HandlerDB) AuthorizationUser(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Debug("got request with bad method", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("Content-Type") != "application/json" {
			log.Debug("wrong Content-Type", req.Header.Get("Content-Type"))
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// десериализуем запрос в структуру модели
		log.Debug("decoding request")
		var jsonUsers models.UserData
		dec := json.NewDecoder(req.Body)
		if err := dec.Decode(&jsonUsers); err != nil {
			log.Debug("cannot decode request JSON body", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonUsers.PasswordHash = auth.HashPassword(jsonUsers.Password)
		err := m.Storage.GetUser(ctx, &jsonUsers)
		if err != nil {
			log.Debug("failed to autorization", err)
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		err = auth.CreateJwtToken(&jsonUsers)
		if err != nil {
			log.Debug("token not created")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		m.DataJWT.AddToken(&jsonUsers)
		res.Header().Add(models.HeaderHTTP, jsonUsers.Token)
		res.WriteHeader(http.StatusOK)

	}

}
