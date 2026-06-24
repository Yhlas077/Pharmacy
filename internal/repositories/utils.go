package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// CONNECT DB
func ConnectDB(config string) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DB error: %v\n", err)
		os.Exit(1)
	}
	db = conn
}

func GetDB() *pgxpool.Pool {
	return db
}
