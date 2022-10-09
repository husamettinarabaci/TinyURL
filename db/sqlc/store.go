package db

import (
	"context"
	"database/sql"
	"errors"
)

type Store interface {
	Querier
	CreateTransformTx(ctx context.Context, arg CreateTransformParams) (Transform, error)
	DeleteTransformTx(ctx context.Context, transformID int64) error
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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

func (store *SQLStore) CreateTransformTx(ctx context.Context, arg CreateTransformParams) (Transform, error) {
	var transform Transform
	err := store.execTx(ctx, func(q *Queries) error {

		user, err := q.GetUser(ctx, arg.Owner)
		if err != nil {
			return err
		}

		if user.UrlCount > 0 && user.UserType == "FREE" {
			return errors.New("you have reached your limit")
		}

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

func (store *SQLStore) DeleteTransformTx(ctx context.Context, transformID int64) error {
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
