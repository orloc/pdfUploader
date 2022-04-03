package repository

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getQueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func execQuery(ctx context.Context, pool *pgxpool.Pool, sql string, args interface{}, handler func(rows pgx.Rows) (interface{}, error)) (interface{}, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var (
		qErr error
		rows pgx.Rows
	)

	if args != nil {
		var arg interface{}
		arr, ok := args.([]interface{})
		if ok == false {
			return nil, errors.New("cannot cast args")
		}
		if len(arr) == 1 {
			arg = arr[0]
			rows, qErr = conn.Query(ctx, sql, arg)
		} else {
			rows, qErr = conn.Query(ctx, sql, args.([]interface{})...)
		}
	} else {
		rows, qErr = conn.Query(ctx, sql)
	}

	if qErr != nil {
		return nil, qErr
	}

	if handler != nil {
		return handler(rows)
	}

	return rows, nil
}
