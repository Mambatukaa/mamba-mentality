package internal

import (
	"log"
	"net/http"
	"os"

	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"mamba-mentality.com/graph"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	if DATABASE_URL == "" {
		fmt.Fprintf(os.Stderr, "Make sure your DB connection url okay? URL: %v", DATABASE_URL)
		os.Exit(1)
	}

	// connect PostreSQL
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("DB connection successfully. 🚀")
	defer dbpool.Close()

}

func NewServer() string {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	handler := "Auth"

	return handler
}
