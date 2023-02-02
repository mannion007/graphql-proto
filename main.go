package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Transaction struct {
	Merchant    string `json:"merchant"`
	Description string `json:"description"`
}

func main() {
	// Define the Transaction object type
	transactionType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"merchant": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// Define the RootQuery fields
	fields := graphql.Fields{
		"transactions": &graphql.Field{
			Type: graphql.NewList(transactionType),
			Args: graphql.FieldConfigArgument{
				"offset": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				transactions := []Transaction{
					{Merchant: "Apple", Description: "iPhone 10"},
					{Merchant: "Google", Description: "Pixel 5"},
					{Merchant: "Samsung", Description: "Galaxy S21"},
					{Merchant: "Microsoft", Description: "Surface Pro 7"},
					{Merchant: "Amazon", Description: "Kindle Paperwhite"},
					{Merchant: "Apple", Description: "MacBook Pro 16"},
					{Merchant: "Google", Description: "Nest Hub Max"},
					{Merchant: "Samsung", Description: "Galaxy Watch 3"},
					{Merchant: "Microsoft", Description: "Xbox Series X"},
					{Merchant: "Amazon", Description: "Echo Dot (4th Gen)"},
					{Merchant: "Apple", Description: "AirPods Max"},
					{Merchant: "Google", Description: "Chromebook Pixelbook Go"},
					{Merchant: "Samsung", Description: "Galaxy Tab S7"},
					{Merchant: "Microsoft", Description: "Surface Laptop Go"},
					{Merchant: "Amazon", Description: "Fire TV Stick Lite"},
					{Merchant: "Apple", Description: "iPad Pro 11"},
					{Merchant: "Google", Description: "Pixel Buds"},
					{Merchant: "Samsung", Description: "Galaxy Z Fold 2"},
					{Merchant: "Microsoft", Description: "Surface Headphones 2"},
					{Merchant: "Amazon", Description: "Echo Show 8"},
				}

				offset := p.Args["offset"].(int)
				limit := p.Args["limit"].(int)

				// Return the sliced transactions based on the offset and limit
				return transactions[offset : offset+limit], nil
			},
		},
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: fields}),
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
