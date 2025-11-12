package db

import (
	"chat-service/config"
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Connection struct {
	conn *pgx.Conn
	ctx  context.Context
}

//go:embed migrations/1_create_base_tables.sql
var createBaseTables string

var migrations = []string{
	createBaseTables,
}

func NewConn(ctx context.Context, cfg config.PostgresConfig) (connection *Connection, err error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHostname, cfg.PostgresPort, cfg.PostgresDatabase, cfg.PostgresSslmode)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	out := Connection{conn: conn, ctx: ctx}
	if _, err := conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS schema_version(version integer not null primary key)"); err != nil {
		out.Close()
		return &out, err
	}
	out.runMigrations(ctx)
	return &out, nil
}

func (conn *Connection) runMigrations(ctx context.Context) error {
	for i, migration := range migrations {
		if _, err := conn.Exec(migration); err != nil {
			conn.Close()
			return err
		}
		conn.Exec("INSERT OR REPLACE INTO schema_version (version) VALUES (?)", i+1)
	}
	return nil
}

func (conn *Connection) Close() {
	conn.conn.Close(conn.ctx)
}

func (conn *Connection) QueryRow(query string, args ...any) pgx.Row {
	return conn.conn.QueryRow(conn.ctx, query, args...)
}

func (conn *Connection) Query(query string, args ...any) (rows pgx.Rows, err error) {
	return conn.conn.Query(conn.ctx, query, args...)
}

func (conn *Connection) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	return conn.conn.Exec(conn.ctx, query, args...)
}

func (conn *Connection) CreateRow(table string, args ...any) (pgconn.CommandTag, error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES (%s);", table, prepareValues(args...))
	return conn.Exec(query, args...)
}

func (conn *Connection) SelectAll(table string) (rows pgx.Rows, err error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	return conn.Query(query)
}

func prepareValues(args ...any) (values string) {
	for i, _ := range args {
		values += fmt.Sprintf("$%d", i+1)
		if i != len(args)-1 {
			values += ", "
		}
	}
	return
}
