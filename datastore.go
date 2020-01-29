// A minimal interface to expose datastore related functions.
// Author: Tiago Melo (tiagoharris@gmail.com)

package datastore

import (
	"context"
)

type Datastore interface {
	GetFeatureFlagByName(ctx context.Context, name string) (*FeatureFlag, error)
}
