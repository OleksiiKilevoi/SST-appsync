package main

import (
	"context"
	"fmt"

	"my-sst-app/graphql/model"
	// "my-sst-app/graphql/resolvers"

	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/mefellows/vesper"
	// "github.com/pkg/errors"
)

// type Event = resolvers.Args[struct {
// 	// Filters *model.TransactionFilters `json:"filters"`
// }]

func handler(ctx context.Context) (*model.User, error) {
	fmt.Println("In handler")
// 	r, err := resolvers.NewResolver()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Resolver:", r)
// 	fmt.Println("ctx:", ctx)
// 	// user, err := 
// 	// if err != nil {
// 		// return nil, errors.Wrap(err, "Can't get user")
// 	// }

	// fmt.Println("User:", user)
// 	return r.Query().Users(ctx)
return nil, nil
}

func main() {
	fmt.Println("here")
	lambda.Start(handler)
	// fmt.Println("Starting lambda...")
	// vesper.New(handler,
	// 	vesper.JSONParserMiddleware(),
	// ).
	// 	DisableAutoUnmarshal().
	// 	Start()
}
