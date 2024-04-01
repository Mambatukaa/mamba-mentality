package dbModels

import (
	"context"
	"fmt"
	"os"
	"time"

	"mamba-mentality.com/internal/db"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string

	CreatedAt time.Time
	Password  string
}

func CreateUser(ctx context.Context) error {
	db := db.GetDBPool()

	var id int
	var title string

	err := db.QueryRow(ctx, "SELECT id, title FROM products WHERE id=1").Scan(&id, &title)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, title)

	return nil
}
