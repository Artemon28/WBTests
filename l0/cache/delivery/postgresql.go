package delivery

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

func (r *repository) Create(ctx context.Context, delivery *cacheModel.Delivery) error {
	q := `
		INSERT INTO delivery 
		    (name, phone_number, zip, city, address, region, email) 
		VALUES 
		       ($1, $2, $3, $4, $5, $6, $7)
	`
	var createString string
	if err := r.client.QueryRow(ctx, q, delivery.Name, delivery.Phone, delivery.Zip,
		delivery.City, delivery.Address, delivery.Region, delivery.Email).
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

func (r *repository) FindOne(ctx context.Context, phone string) (cacheModel.Delivery, error) {
	q := `
		SELECT name, phone_number, zip, city, address, region, email FROM public.delivery WHERE phone_number = $1
	`
	var del cacheModel.Delivery
	err := r.client.QueryRow(ctx, q, phone).Scan(&del.Name, &del.Phone, &del.Zip, &del.City, &del.Address, &del.Region, &del.Email)
	if err != nil {
		return cacheModel.Delivery{}, err
	}

	return del, nil
}

func NewRepository(client postgresDataBase.Client) *repository {
	return &repository{
		client: client,
	}
}
