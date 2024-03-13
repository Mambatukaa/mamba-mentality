package main

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
	"mamba-mentality.com/internal/graphql"
)

func init() {
	err := godotenv.Load()
	
  if err != nil {
		fmt.Print("hahah", err)
    log.Fatal("Error loading .env file")
  }

	DATABASE_URL := os.Getenv("DATABASE_URL");

	if(DATABASE_URL == "") {
		fmt.Fprintf(os.Stderr, "Make sure your DB connection url okay? URL: %v", DATABASE_URL);
		os.Exit(1);
	}

	// connect PostreSQL
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("DB connection successfully. 🚀")
	defer dbpool.Close()

};

func main() {
	PORT := os.Getenv("PORT")

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground 🚀🚀🚀", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
