package database

import (
	"context"
	"database/sql"
)

// ConnectionsPool define uma interface para connections pool
type ConnectionsPool interface {
	NewTransaction(ctx context.Context, readOnly bool) (tx Transaction, err error)
}

// Transaction define uma interface de transação
type Transaction interface {
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
	Commit() error
	Rollback()
}
