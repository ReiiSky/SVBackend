package databases

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConnectionConfig struct {
	Hostname string
	Port     int
	DBname   string
	Username string
	Password string
	Queries  string
}

type databaseConnection struct {
	db *sqlx.DB
}

func NewDatabaseConnection(config DatabaseConnectionConfig) (*databaseConnection, error) {
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.Username, config.Password, config.Hostname, config.Port, config.DBname, config.Queries)

	db, err := sqlx.Open("mysql", connectionURL)

	if err != nil {
		return nil, err
	}

	dc := &databaseConnection{
		db,
	}

	if err := dc.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dc, nil
}

func (dc databaseConnection) Ping(ctx context.Context) error {
	return dc.db.PingContext(ctx)
}

func (dc databaseConnection) RunQuery(ctx context.Context, result any, query string, args ...any) error {
	return dc.db.SelectContext(ctx, result, query, args...)
}

func (dc databaseConnection) RunCommand(ctx context.Context, operationFn func(*sqlx.Tx) error) error {
	tx := dc.db.MustBeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})

	isCommited := false

	defer func() {
		if isCommited {
			return
		}

		tx.Rollback()
	}()

	if err := operationFn(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	isCommited = true
	return nil
}
