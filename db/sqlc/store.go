package db

import (
	"context"
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

type TransformTxParams struct {
	Owner    string `json:"owner"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func (store *Store) CreateTransformTx(ctx context.Context, arg TransformTxParams) (Transform, error) {
	var transform Transform
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		transform, err = q.CreateTransform(ctx, CreateTransformParams(arg))
		if err != nil {
			return err
		}

		_, err = q.IncUrlCount(ctx, arg.Owner)
		if err != nil {
			return err
		}

		return nil
	})
	return transform, err
}

func (store *Store) DeleteTransformTx(ctx context.Context, transformID int64) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		transform, err := q.GetTransform(ctx, transformID)
		if err != nil {
			return err
		}

		err = q.DeleteTransform(ctx, transform.ID)
		if err != nil {
			return err
		}

		_, err = q.DecUrlCount(ctx, transform.Owner)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
