package storage

import (
	"context"
	"time"

	"github.com/MlDenis/internal/gofermart/models"
	log "github.com/sirupsen/logrus"
	"github.com/MlDenis/pkg"
)

// записываем данные нового пользователя в бд
func (pgdb *PostgresDB) LoadOrderInDB(ctx context.Context, orders *models.Orders) error {
	orders.OrderDate = time.Now()
	orders.StatusOrder = models.NewOrder
	tx, err := pgdb.pool.Begin(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO public.orders (ordernumber,userlogin,orderdate,statusorder) VALUES ($1, $2,$3, $4)`,
		orders.OrderNumber, orders.UserLogin, orders.OrderDate, orders.StatusOrder,
	)
	if err != nil {
		log.Error(err)
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)

}

func (pgdb *PostgresDB) GetUserOrders(ctx context.Context, user *models.UserData) ([]models.Orders, error) {
	orders := []models.Orders{}
	rows, err := pgdb.pool.Query(ctx, `SELECT * FROM orders WHERE userlogin = $1`, user.Login)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		order := models.Orders{}
		rows.Scan(&order)
		orders = append(orders, order)
	}

	defer rows.Close()

	if len(orders) == 0 {
		return orders, pkg.NoOrders
	}

	return orders, nil
}
