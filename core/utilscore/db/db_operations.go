package dbutils

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)


func SetterLimitAndOffsetQuery(query squirrel.SelectBuilder, offset *uint64, limit *uint64) squirrel.SelectBuilder {
	if limit != nil {
		query = query.Limit(*limit)
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	return query
}

func RollbackTransactionDB(ctx context.Context, tx pgx.Tx) {
	if rErr := tx.Rollback(ctx); rErr != nil && !errors.Is(rErr, pgx.ErrTxClosed) {
		log.Printf("error: failed to rollback transaction: %v", rErr)
	}
}

func BeginTransaction(ctx context.Context, databasePull *pgxpool.Pool) (*pgxpool.Conn, pgx.Tx, error) {
	conn, err := databasePull.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		conn.Release()
		return nil, nil, err
	}
	return conn, tx, nil
}
