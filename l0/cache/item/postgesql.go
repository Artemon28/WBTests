package item

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

func (r *repository) Create(ctx context.Context, item *cacheModel.Item) error {
	q := `
		INSERT INTO item 
		    (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
		VALUES 
		       ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	var createString string
	if err := r.client.QueryRow(ctx, q, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale,
		item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status).
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

func (r *repository) FindAll(ctx context.Context, order_uid string) ([]cacheModel.Item, error) {
	q := `SELECT public.item.chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
			FROM public.item
			JOIN public.items_in_order ON public.item.chrt_id=public.items_in_order.chrt_id
			WHERE public.items_in_order.order_uid = $1
	`

	items, err := r.client.Query(ctx, q, order_uid)
	var itemsInOrder = []cacheModel.Item{}
	if err != nil {
		return nil, err
	}
	for items.Next() {
		var item cacheModel.Item

		err = items.Scan(&item.Chrt_id, &item.Track_number, &item.Price, &item.Rid, &item.Name, &item.Sale,
			&item.Size, &item.Total_price, &item.Nm_id, &item.Brand, &item.Status)
		if err != nil {
			return nil, err
		}

		itemsInOrder = append(itemsInOrder, item)
	}

	if err = items.Err(); err != nil {
		return nil, err
	}

	return itemsInOrder, nil
}

func NewRepository(client postgresDataBase.Client) *repository {
	return &repository{
		client: client,
	}
}
