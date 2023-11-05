package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"spark/utils"
)

func OpenConnection() (pool ConnectionsPool, err error) {
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
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &connections{db: db}, nil
}

// NewTransaction define um metodo para criar uma nova transação no
// banco de dados
func NewTransaction(ctx context.Context, readOnly bool) (tx Transaction, err error) {
	conn, ok := ctx.Value(utils.DatabaseKey).(ConnectionsPool)
	if !ok {
		return nil, errors.New("erro ao extrair a connctions pool")
	}
	return conn.NewTransaction(ctx, readOnly)
}
