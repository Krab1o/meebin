package transaction

import (
	"context"

	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/Krab1o/meebin/internal/client/db/pg"
	"github.com/jackc/pgx/v5"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) wrapToTransaction(
	ctx context.Context,
	opts pgx.TxOptions,
	fn db.Handler,
) error {
	// Проверяем, была ли уже у нас внешняя транзакция
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		//Если была, то обрачиваем хендлер в эту транзакцию
		return fn(ctx)
	}

	// Если не было, создаём новую
	tx, err := m.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	// Затем помещаем в контекст
	ctx = context.WithValue(ctx, pg.TxKey, tx)

	defer func() {
		// something bad happened during transaction
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				//rollback failed
			}
			return
		}

		// transaction happened successfully
		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				//tx commit failed
			}
		}
	}()

	// И выполняем действия с новосозданной транзакцией
	if err = fn(ctx); err != nil {
		return err
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, fn db.Handler) error {
	txOptions := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.wrapToTransaction(ctx, txOptions, fn)
}
