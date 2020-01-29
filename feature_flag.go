// Author: Tiago Melo (tiagoharris@gmail.com)

package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
)

// FeatureFlag is used to store information about a feature flag for a given system
type FeatureFlag struct {
	Name     string `datastore:"name" json:"name"`
	IsActive bool   `datastore:"is_active" json:"is_active"`
}

// FeatureFlagKey creates a new datastore key for a given entity type and feature flag name
func FeatureFlagKey(ctx context.Context, kind, name string) *datastore.Key {
	return datastore.NameKey(kind, name, nil)
}

// GetFeatureFlagByName queries datastore for a given entity type and feature flag name
func GetFeatureFlagByName(ctx context.Context, s *Store, name string) (*FeatureFlag, error) {
	var f FeatureFlag
	err := s.Client.Get(ctx, FeatureFlagKey(ctx, s.Kind, name), &f)

	return &f, err
}

// GetFeatureFlagByName returns the feature flag of a given system
func (s *Store) GetFeatureFlagByName(ctx context.Context, name string) (*FeatureFlag, error) {
	flag, err := GetFeatureFlagByName(ctx, s, name)
	return flag, err
}
