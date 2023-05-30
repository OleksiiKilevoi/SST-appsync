package account

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"

	"my-sst-app/dao"
)

type Controller struct {
	Sess  *session.Session
	Users *dynamo.Table
	DDB   *dynamo.DB
}

var ErrAccountNotFound = errors.New("account not found")
var ErrAccountIDExists = errors.New("account ID already exists")
var ErrAccountEmailExists = errors.New("email already exists")
var ErrAccountPhoneNumberExists = errors.New("phone number already exists")

type NewControllerConfig struct {
	Session *session.Session
}

func NewController(config *NewControllerConfig) (*Controller, error) {
	ddbClient := dynamo.New(config.Session)
	users := ddbClient.Table("users")

	return &Controller{
		Sess:  config.Session,
		Users: &users,
		DDB:   ddbClient,
	}, nil
}

func (c *Controller) GetAccountByID(ctx context.Context, id string) (*dao.User, error) {
	user := new(dao.User)
	fmt.Printf("here")
	if err := c.Users.
		Get("id", id).
		OneWithContext(ctx, user); err != nil {
		if err == dynamo.ErrNotFound {
			return nil, ErrAccountNotFound
		}
		return nil, errors.Wrap(err, "failed to get item from dynamo")
	}

	return user, nil
}
