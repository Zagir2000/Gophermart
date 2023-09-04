package handlers

import (
	"context"

	"github.com/MlDenis/internal/gofermart/handlers/balance"
	"github.com/MlDenis/internal/gofermart/handlers/order"
	"github.com/MlDenis/internal/gofermart/handlers/users"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, newHandStruct *HandlerDB) chi.Router {
	Balance := balance.HandlerBalance(newHandStruct.Storage, newHandStruct.DataJWT)
	Users := users.HandlerUsers(newHandStruct.Storage, newHandStruct.DataJWT)
	Orders := order.HandlerOrders(newHandStruct.Storage, newHandStruct.DataJWT)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/api/user/register", Users.RegisterNewUser(ctx))
	r.Post("/api/user/login", Users.AuthorizationUser(ctx))
	r.Post("/api/user/orders", Orders.LoadOrderNumber(ctx))
	r.Get("/api/user/orders", Orders.GetUserOrder(ctx))
	r.Get("/api/user/balance", Balance.GetBalance(ctx))
	r.Post("/api/user/balance/withdraw", Balance.WithdrawBalance(ctx))
	r.Get("/api/user/withdrawals", Balance.GetWithdrawals(ctx))
	return r
}
