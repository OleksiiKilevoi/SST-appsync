package main

import (
	"context"
	"fmt"

	"my-sst-app/graphql/model"
	"my-sst-app/graphql/resolvers"

	"github.com/mefellows/vesper"
	"github.com/pkg/errors"
)

func handler(ctx context.Context) (*model.User, error) {
	fmt.Println("In handler")
	r, err := resolvers.NewResolver()
	if err != nil {
		return nil, err
	}

	fmt.Println("Resolver:", r)
	fmt.Println("ctx:", ctx)
	user, err := r.Query().Users(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Can't get user")
	}

	fmt.Println("User:", user)
	return user, nil
}

func main() {
	fmt.Println("Starting lambda...")
	vesper.New(handler,
		vesper.JSONParserMiddleware(),
	).
		DisableAutoUnmarshal().
		Start()
}
