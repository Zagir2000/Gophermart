package storage

import (
	"context"

	"github.com/MlDenis/internal/accrual/models"
	log "github.com/sirupsen/logrus"
)

func (pgdb *PostgresDB) AddGoods(ctx context.Context, orderForRegister *models.OrderForRegister) error {
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, goods := range orderForRegister.Goods {
		_, err = tx.Exec(ctx,
			`INSERT INTO public.goods (descriptionorder,price) VALUES ($1, $2)`,
			goods.Description, goods.Price,
		)
		if err != nil {
			log.Error(err)
			tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}
func (pgdb *PostgresDB) RegisterInfoInDB(ctx context.Context, goods *models.Reward) error {
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	match := "%" + goods.Match + "%"
	_, err = tx.Exec(ctx,
		`UPDATE public.goods set reward = $1, rewardtype = $2 WHERE descriptionorder like $3`,
		goods.Reward, goods.RewardType, match,
	)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}
