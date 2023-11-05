package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"spark/utils"
)

func OpenConnection() (db *sql.DB, err error) {
	var connString string = fmt.Sprintf("host=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable",
		viper.GetString("database_host"),
		viper.GetString("database_user"),
		viper.GetString("database_password"),
		viper.GetString("database_name"),
	)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewTransaction define um metodo para criar uma nova transação no
// banco de dados
func NewTransaction(ctx context.Context) (tx *sql.Tx, err error) {
	db := ctx.Value(utils.DatabaseKey).(*sql.DB)

	if tx, err = db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	}); err != nil {
		return
	}

	return
}
