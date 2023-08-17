package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/MlDenis/internal/gofermart/models"
	"github.com/MlDenis/internal/luna"
	"github.com/MlDenis/pkg"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)

// авторизация
func (m *HandlerDB) LoadOrderNumber(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Debug("got request with bad method", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("Content-Type") != "text/plain" {
			log.Debug("wrong Content-Type", req.Header.Get("Content-Type"))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonUsers := &models.UserData{}
		jsonOrders := &models.Orders{}
		jsonUsers.Token = req.Header.Get(models.HeaderHTTP)
		err := m.DataJWT.GetToken(jsonUsers)
		if err != nil {
			log.Debug("user not authenticated", err)
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		number, err := io.ReadAll(req.Body)
		if err != nil {
			log.Debug(err)
			res.WriteHeader(http.StatusBadRequest)
		}

		orderID, err := strconv.ParseInt(string(number), 10, 64)
		validNumber := luna.Valid(orderID)
		if validNumber == false {
			log.Debug("invalid order number")
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		jsonOrders.OrderNumber = orderID
		jsonOrders.UserLogin = jsonUsers.Login
		err = m.pgDB.LoadOrderInDB(ctx, jsonOrders)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == pkg.UniqueViolationCode {
					log.Debug("the order number has already been uploaded by the user")
					res.WriteHeader(http.StatusOK)
					return
				}
				log.Debug("error in add orders in db", err)
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Debug("error in add orders in db", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Debug("new order number accepted for processing")
		res.WriteHeader(http.StatusAccepted)
		return
	}

}
