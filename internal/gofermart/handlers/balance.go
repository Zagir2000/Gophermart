package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MlDenis/internal/gofermart/models"
	"github.com/MlDenis/internal/luna"
	"github.com/MlDenis/pkg"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)

// баланс пользователя
func (m *HandlerDB) GetBalance(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			log.Error("got request with bad method", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		jsonUsers := &models.UserData{}
		jsonUsers.Token = req.Header.Get(models.HeaderHTTP)
		err := m.DataJWT.GetToken(jsonUsers)
		if err != nil {
			log.Error("user not authenticated: ", err)
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		ResponseBalance, err := m.Storage.GetBalanceDB(ctx, jsonUsers.Login)
		if err != nil {
			log.Error("error in add orders in db: ", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		response, err := json.Marshal(ResponseBalance)
		if err != nil {
			log.Error("cannot marshal to json: ", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(response)
		return
	}
}

// Запрос на списание средств
func (m *HandlerDB) WithdrawBalance(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			log.Error("got request with bad method: ", req.Method)
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("Content-Type") != "application/json" {
			log.Error("wrong Content-Type", req.Header.Get("Content-Type"))
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		//Создадим структуру пользователя, чтобы записать токен и логин в него
		jsonUsers := &models.UserData{}
		jsonUsers.Token = req.Header.Get(models.HeaderHTTP)
		err := m.DataJWT.GetToken(jsonUsers)
		if err != nil {
			log.Error("user not authenticated: ", err)
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		wisthdrawSum := &models.WithdrawOrder{}
		// десериализуем запрос в структуру модели
		log.Error("decoding request")
		dec := json.NewDecoder(req.Body)
		if err := dec.Decode(&wisthdrawSum); err != nil {
			log.Error("cannot decode request JSON body", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		//проверям заказ через алгоритм луна
		validNumber := luna.Valid(wisthdrawSum.Order)
		if !validNumber {
			log.Error("invalid order number:")
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		//Проверям баланс
		ResponseBalance, err := m.Storage.GetBalanceDB(ctx, jsonUsers.Login)
		if err != nil {
			log.Error("error in add orders in db: ", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		//проверям хватает ли средств для списания
		if ResponseBalance.AccrualSum < wisthdrawSum.Sum {
			log.Error("insufficient funds to write off: ", err)
			res.WriteHeader(http.StatusPaymentRequired)
			return
		}
		//Создадим структуру заказов, чтобы записать их в бд
		jsonOrders := &models.Orders{}
		jsonOrders.OrderNumber = wisthdrawSum.Order
		jsonOrders.UserLogin = jsonUsers.Login
		jsonOrders.StatusOrder = models.WithdrawEnd
		jsonOrders.Withdraw = wisthdrawSum.Sum
		err = m.Storage.LoadOrderInDB(ctx, jsonOrders)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				//Проверяем,не загружали ли этот заказ ранее(точно ли это ноывй заказ)
				if pgErr.Code == pkg.UniqueViolationCode {
					log.Error("cannot be debited for this order because this order has already been uploaded")
					res.WriteHeader(http.StatusConflict)
					return
				}
				log.Error("error in add orders in db: ", err)
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Error("error in add orders in db: ", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		//Изменим баланс после списания
		err = m.Storage.EditBalanceWithdraw(ctx, jsonOrders.UserLogin, wisthdrawSum.Sum)
		if err != nil {
			log.Error("balance change error: ", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(ResponseBalance)
		if err != nil {
			log.Error("cannot make json balance: ", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(response)
		res.WriteHeader(http.StatusOK)
		return
	}
}
