package payment

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

func (r *repository) Create(ctx context.Context, payment *cacheModel.Payment) error {
	q := `
		INSERT INTO payment 
		    (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
		VALUES 
		       ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	var createString string
	if err := r.client.QueryRow(ctx, q, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider,
		payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee).
		Scan(&createString); err != nil {
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

func (r *repository) FindOne(ctx context.Context, trans string) (cacheModel.Payment, error) {
	q := `
		SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total,
		custom_fee FROM public.payment WHERE transaction = $1
	`
	var pay = cacheModel.Payment{}
	err := r.client.QueryRow(ctx, q, trans).Scan(&pay.Transaction, &pay.Request_id, &pay.Currency, &pay.Provider, &pay.Amount, &pay.Payment_dt,
		&pay.Bank, &pay.Delivery_cost, &pay.Goods_total, &pay.Custom_fee)
	if err != nil {
		return cacheModel.Payment{}, err
	}

	return pay, nil
}

func NewRepository(client postgresDataBase.Client) *repository {
	return &repository{
		client: client,
	}
}
