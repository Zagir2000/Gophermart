package handlers

import (
	"context"

	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, newHandStruct *HandlerDB) chi.Router {
	r := chi.NewRouter()
	r.Post("/api/user/register", newHandStruct.RegisterNewUser(ctx))
	r.Post("/api/user/login", newHandStruct.AuthorizationUser(ctx))
	r.Post("/api/user/orders", newHandStruct.LoadOrderNumber(ctx))
	return r
}
