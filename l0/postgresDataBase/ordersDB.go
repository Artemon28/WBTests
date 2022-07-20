package postgresDataBase

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type ConnectionConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cc ConnectionConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cc.Username, cc.Password, cc.Host, cc.Port, cc.Database)
	pool, err = pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, nil
}
