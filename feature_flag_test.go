// Author: Tiago Melo (tiagoharris@gmail.com)

package datastore

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestGetFeatureFlagByName(t *testing.T) {
	ctx := context.Background()
	testKind := "test-feature-flags"

	store := newTestDB(ctx, testKind).(*Store)

	featureFlagName := "Test Flag"
	featureFlag, err := store.GetFeatureFlagByName(ctx, featureFlagName)
	if err == nil {
		t.Error("expected 'datastore.ErrNoSuchEntity' error")
	}

	f := FeatureFlag{
		Name:     featureFlagName,
		IsActive: true,
	}
	_, err = store.Client.Put(ctx, datastore.NameKey(testKind, featureFlagName, nil), &f)
	if err != nil {
		t.Errorf("Creating feature flag entry %s", err)
	}
	featureFlag, err = store.GetFeatureFlagByName(ctx, featureFlagName)
	if err != nil {
		t.Errorf("GetFeatureFlagByName %s", err)
	}
	if !featureFlag.IsActive {
		t.Errorf("Expected feature flag to be active, got %v", featureFlag.IsActive)
	}
}
