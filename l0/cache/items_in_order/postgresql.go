package items_in_order

import (
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

func (r *repository) Create(ctx context.Context, order_uid string, chrt_id int) error {
	q := `
		INSERT INTO items_in_order
		    (order_uid, chrt_id) 
		VALUES 
		       ($1, $2)
	`
	var createString string
	if err := r.client.QueryRow(ctx, q, order_uid, chrt_id).
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

func NewRepository(client postgresDataBase.Client) *repository {
	return &repository{
		client: client,
	}
}
