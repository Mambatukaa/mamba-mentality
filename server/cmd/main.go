package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"mamba-mentality.com/internal/db"
	"mamba-mentality.com/internal/graphql"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db connection
	db.DbConnection()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	PORT := os.Getenv("PORT")

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground 🚀🚀🚀", PORT)

	r.Run()
	db.CloseDBPool()

}

// func main() {
// 	PORT := os.Getenv("PORT")

// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground 🚀🚀🚀", PORT)
// 	log.Fatal(http.ListenAndServe(":"+PORT, nil))

// 	// close db pool
// 	config.CloseDBPool()
// }
