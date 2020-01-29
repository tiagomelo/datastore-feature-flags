// Author: Tiago Melo (tiagoharris@gmail.com)

package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

func newTestDB(ctx context.Context, kind string) Datastore {
	dsClient, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		panic(fmt.Sprintf("could not create new datastore client: %s", err))
	}

	return &Store{
		Client: dsClient,
		Kind:   kind,
	}
}
