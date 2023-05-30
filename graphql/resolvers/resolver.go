package resolvers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"

	users "my-sst-app/controllers/users"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Args[TArgs any] struct {
	Arguments TArgs `json:"arguments"`
}

type Resolver struct {
	Sess      *session.Session
	UsersCtrl *users.Controller
}

func NewResolver() (*Resolver, error) {
	sess, err := session.NewSession()
	fmt.Println(sess)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AWS session")
	}
	usersCtr, _ := users.NewController(&users.NewControllerConfig{
		Session: sess,
	})

	return &Resolver{
		Sess:      sess,
		UsersCtrl: usersCtr,
	}, nil
}

var ErrVerifyCodeInvalid = errors.New("invalid verification code")
