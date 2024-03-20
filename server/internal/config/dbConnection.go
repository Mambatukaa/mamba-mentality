package config

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbpool *pgxpool.Pool
	dbmu   sync.Mutex
)

func DbConnection() {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	ctx := context.Background()

	if DATABASE_URL == "" {
		fmt.Fprintf(os.Stderr, "Make sure your DB connection url okay? URL: %v", DATABASE_URL)
		os.Exit(1)
	}

	// connect PostreSQL
	var err error
	dbpool, err = pgxpool.New(ctx, DATABASE_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("DB connection successfully. 🚀")
}

func GetDBPool() *pgxpool.Pool {
	dbmu.Lock()
	defer dbmu.Unlock()
	return dbpool
}

func CloseDBPool() {
	dbmu.Lock()
	defer dbmu.Unlock()
	if dbpool != nil {
		dbpool.Close()
	}
}
