// Author: Tiago Melo (tiagoharris@gmail.com)

package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

type Store struct {
	Client *datastore.Client

	// Kind represents the entity type to be queried
	Kind string
}

// NewStore initializes a new datastore client
func NewStore(ctx context.Context, kind string) (Datastore, error) {
	dsClient, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to init datastore client")
	}

	return &Store{
		Client: dsClient,
		Kind:   kind,
	}, nil
}
