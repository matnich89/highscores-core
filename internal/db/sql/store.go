package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
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
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) CharacterCreationTx(ctx context.Context, arg CreateCharacterParams) (*Character, error) {
	var result Character

	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateCharacter(ctx, arg)
		result = res
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (store *Store) CharacterUpdateTx(ctx context.Context, arg UpdateCharacterParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		err := q.UpdateCharacter(ctx, arg)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (store *Store) XpRecordCreationTx(ctx context.Context, arg CreateXpRecordParams) (*XpRecord, error) {
	var result *XpRecord

	err := store.execTx(ctx, func(q *Queries) error {
		res, err := q.CreateXpRecord(ctx, arg)
		result = &res
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
