package storage

import (
	"context"

	"github.com/MlDenis/internal/gofermart/models"
)

// записываем данные нового пользователя в бд
func (pgdb *PostgresDB) RegisterUser(ctx context.Context, userData models.UserData) error {

	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {

		return err
	}

	_, err = tx.Exec(ctx, `INSERT INTO public.users (userlogin,hashpass) VALUES ($1, $2)`, userData.Login, userData.PasswordHash)
	if err != nil {

		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)

}

// проверям есть ли пользователь в бд
func (pgdb *PostgresDB) GetUser(ctx context.Context, userData *models.UserData) error {
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {

		return err
	}
	var userId int64
	row := pgdb.pool.QueryRow(ctx, "SELECT id FROM public.users WHERE userlogin=$1", userData.Login)
	err = row.Scan(&userId)
	if err != nil {

		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}
