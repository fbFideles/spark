package database

import "database/sql"

type transaction struct {
	tx *sql.Tx
}

func (t *transaction) Exec(query string, args ...any) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *transaction) QueryRow(query string, args ...any) *sql.Row {
	return t.tx.QueryRow(query, args...)
}

func (t *transaction) Commit() error {
	return t.tx.Commit()
}
func (t *transaction) Rollback() {
	_ = t.tx.Rollback()
}
