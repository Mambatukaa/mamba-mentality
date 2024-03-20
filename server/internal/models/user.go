package dbModels

import (
	"context"
	"fmt"
	"os"

	"mamba-mentality.com/internal/config"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func CreateUser(ctx context.Context) error {
	db := config.GetDBPool()

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
