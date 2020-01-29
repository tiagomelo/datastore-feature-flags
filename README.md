# datastore-feature-flags

A small toolkit that enable any Golang project to use Datastore entities (https://cloud.google.com/datastore/) as a simple feature flags mechanism (https://en.wikipedia.org/wiki/Feature_toggle).

### Basic usage

Given the following entity on GCP Datastore:

![](images/datastore.png)

This is a sample code to retrieve it:

```
package main

import (
	"context"

	"github.com/tiagomelo/datastore-feature-flags"
)

func main() {
	ctx := context.Background()

	store, err := datastore.NewStore(ctx, "my-api-feature-flags")
	if err != nil {
		log.Fatal(err, "unable to init database")
	}

	featureFlag, err := store.GetFeatureFlagByName(ctx, "Test flag")
	if err != nil {
		// proper error handling
	}

	if featureFlag.IsActive {
		// ...
	}
}
```
