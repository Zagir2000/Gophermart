package storage

import (
	"context"

	"github.com/MlDenis/internal/accrual/models"
	log "github.com/sirupsen/logrus"
)

// Получение баланса пользователя
func (pgdb *PostgresDB) GetOrderFromOrdersAccrualDB(ctx context.Context, ordernumber int64) (*models.Order, error) {
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ordersAccrual := &models.Order{}
	row := pgdb.pool.QueryRow(ctx, `SELECT ordernumber,statusorder,accrual FROM public.ordersaccrual WHERE ordernumber=$1`, ordernumber)
	err = row.Scan(&ordersAccrual.OrderNumber, &ordersAccrual.StatusOrder, &ordersAccrual.Accrual)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
		return nil, err
	}

	return ordersAccrual, tx.Commit(ctx)
}

// Регистрация нового совершённого заказа
func (pgdb *PostgresDB) LoadOrderInOrdersAccrualDB(ctx context.Context, orderForRegister *models.OrderForRegister) error {
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO public.ordersaccrual (ordernumber,statusorder,accrual,goods) VALUES ($1, $2, $3, $4)`,
		orderForRegister.OrderNumber, models.Registered, 0, orderForRegister.Goods,
	)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}
