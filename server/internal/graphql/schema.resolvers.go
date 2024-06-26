package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"

	"mamba-mentality.com/internal/db"
	"mamba-mentality.com/internal/graphql/model"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	db := db.GetDBPool()

	sqlStatement := `
		INSERT INTO users (first_name, last_name)
		VALUES ($1, $2)
		RETURNING id`

	id := 0
	err := db.QueryRow(ctx, sqlStatement, input.FirstName, input.LastName).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)

	firstName := "Batuka"
	lastName := "Nasan"

	return &model.User{
		FirstName: &firstName,
		LastName:  &lastName,
	}, err
}

// AddProduct is the resolver for the addProduct field.
func (r *mutationResolver) AddProduct(ctx context.Context, input *model.ProductInput) (*model.Product, error) {
	db := db.GetDBPool()

	sqlStatement := `
		INSERT INTO products (name, price)
		VALUES ($1, $2)
		RETURNING id`

	id := 0
	err := db.QueryRow(ctx, sqlStatement, input.Name, input.Price).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)

	var name string
	var price string

	err = db.QueryRow(ctx,
		`
		SELECT name, price
		FROM products
		WHERE id=1
		`).Scan(&name, &price)

	if err != nil {
		panic(err)
	}

	fmt.Println(name, price, "hahahahha")

	return &model.Product{
		Name:  &name,
		Price: &price,
	}, err

}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db := db.GetDBPool()
	var (
		firstName string
		lastName  string
	)

	sqlStatement := `
		SELECT first_name, last_name 
		FROM users;
	`

	users, err := db.Query(ctx, sqlStatement)

	if err != nil {
		panic(err)
	}

	defer users.Close()

	for users.Next() {
		err := users.Scan(&firstName, &lastName)

		if err != nil {
			panic(err)
		}

		fmt.Println("\n", firstName, lastName)
	}

	return nil, err
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	db := db.GetDBPool()

	var (
		name  string
		price string
	)

	sqlStatement := `
		SELECT name, price
		FROM products;
	`

	products, err := db.Query(ctx, sqlStatement)

	if err != nil {
		panic(err)
	}

	defer products.Close()

	answer := []*model.Product{}

	for products.Next() {
		err := products.Scan(&name, &price)

		if err != nil {
			panic(err)
		}

		answer = append(answer, &model.Product{Name: &name, Price: &price})

		fmt.Println("\n", name, price)
	}

	return answer, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
