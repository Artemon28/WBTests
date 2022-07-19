package order

import (
	"L0/cacheModel"
	"L0/postgresDataBase"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresDataBase.Client
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, order *cacheModel.Order) error {
	q := `
		INSERT INTO orders 
		    (order_uid, track_number, entry, del_phone, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
		VALUES 
		       ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	var createString string
	if err := r.client.QueryRow(ctx, q, order.Order_uid, order.Track_number, order.Entry,
		order.Delivery.Phone, order.Locale, order.Internal_signature, order.Customer_id,
		order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created,
		order.Oof_shard).Scan(&createString); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (u map[string]cacheModel.Order, err error) {
	q := `
		SELECT order_uid, track_number, entry, del_phone, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM public.orders;
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	orders := make(map[string]cacheModel.Order, 0)

	for rows.Next() {
		var ord cacheModel.Order

		err = rows.Scan(&ord.Order_uid, &ord.Track_number, &ord.Entry, &ord.Del_phone, &ord.Locale,
			&ord.Internal_signature, &ord.Customer_id, &ord.Delivery_service, &ord.Shardkey,
			&ord.Sm_id, &ord.Date_created, &ord.Oof_shard)
		if err != nil {
			return nil, err
		}

		orders[ord.Order_uid] = ord
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func NewRepository(client postgresDataBase.Client) *repository {
	return &repository{
		client: client,
	}
}
