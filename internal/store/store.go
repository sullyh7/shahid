package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

var (
	ErrNotFound  = errors.New("record not found")
	ErrConflict  = errors.New("record already exists")
	QueryTimeout = 5 * time.Second
)

type Storage struct {
	Videos interface {
		Create(context.Context, *Video) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Videos: &VideoStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func isUniqueViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}
	return false
}
