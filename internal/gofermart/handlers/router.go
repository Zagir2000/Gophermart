package handlers

import (
	"context"

	"github.com/go-chi/chi/v5"
)

func Router(ctx context.Context, newHandStruct *MetricHandlerDB) chi.Router {
	r := chi.NewRouter()
	r.Post("/api/user/register", newHandStruct.RegisterNewUser(ctx))
	return r
}
