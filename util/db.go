package util

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var dbPool *pgxpool.Pool
var dbContext context.Context

func AcquirePostgresPool(envUri string) (*pgxpool.Pool, context.Context) {
	if dbPool != nil {
		return dbPool, dbContext
	}

	config, err := pgxpool.ParseConfig(envUri)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 100

	ctx := context.Background()

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	dbPool = pool
	dbContext = ctx
	return dbPool, ctx
}

