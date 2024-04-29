/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vxlan

import (
	"context"
	"errors"
	"fmt"

	"github.com/henderiw/idxtable/pkg/table32"
	"github.com/henderiw/idxtable/pkg/tree"
	"github.com/henderiw/logger/log"
	"github.com/henderiw/store"
	vxlanbev1alpha1 "github.com/kuidio/kuid/apis/backend/vxlan/v1alpha1"
	"k8s.io/utils/ptr"
)

type Applicator interface {
	Validate(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) error
	Apply(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) error
	Delete(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) error
}

type applicator struct {
	cacheCtx *CacheContext
}

func (r *applicator) getEntriesByOwner(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) (map[string]tree.Entries, error) {
	treeEntries := map[string]tree.Entries{}
	ownerSelector, err := claim.GetOwnerSelector()
	if err != nil {
		return nil, err
	}
	claimType := claim.GetClaimType()
	treeEntries[""] = r.cacheCtx.tree.GetByLabel(ownerSelector)
	if len(treeEntries) != 0 {
		// ranges and prefixes using network type can have multiple plrefixes
		if len(treeEntries[""]) > 1 && (claimType != vxlanbev1alpha1.VXLANClaimType_Range) {
			return treeEntries, fmt.Errorf("multiple entries match the owner, %v", treeEntries[""])
		}
	}
	// for id based claims we should also check the range tables
	if claimType != vxlanbev1alpha1.VXLANClaimType_Range {
		var errs error
		r.cacheCtx.ranges.List(ctx, func(ctx context.Context, k store.Key, t *table32.Table32) {
			treeEntries[k.Name] = t.GetByLabel(ownerSelector)
			if len(treeEntries[k.Name]) > 1 {
				errs = errors.Join(errs, fmt.Errorf("multiple entries match the owner, %v", treeEntries[k.Name]))
				return
			}
		})
		if errs != nil {
			return treeEntries, errs
		}
	}
	return treeEntries, nil
}

func (r *applicator) delete(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) error {
	existingEntries, err := r.getEntriesByOwner(ctx, claim)
	if err != nil {
		return err
	}

	for treeName, existingEntries := range existingEntries {
		for _, existingEntry := range existingEntries {
			if treeName == "" {
				r.cacheCtx.tree.ReleaseID(existingEntry.ID())
			} else {
				k := store.ToKey(treeName)
				if table, err := r.cacheCtx.ranges.Get(ctx, k); err == nil {
					if err := table.Release(uint32(existingEntry.ID().ID())); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (r *applicator) getEntriesByLabelSelector(ctx context.Context, claim *vxlanbev1alpha1.VXLANClaim) tree.Entries {
	log := log.FromContext(ctx)
	labelSelector, err := claim.GetLabelSelector()
	if err != nil {
		log.Error("cannot get label selector", "error", err.Error())
		return nil
	}
	return r.cacheCtx.tree.GetByLabel(labelSelector)
}

func (r *applicator) deleteNonClaimedEntries(ctx context.Context, existingEntries map[string]tree.Entries, id *uint32, reclaimTreeName string) error {
	for treeName, existingEntries := range existingEntries {
		fmt.Println("deleteNonClaimedEntries", treeName, existingEntries)
		for _, existingEntry := range existingEntries {
			if id != nil && *id == uint32(existingEntry.ID().ID()) && reclaimTreeName == treeName {
				continue
			}
			if treeName == "" {
				r.cacheCtx.tree.ReleaseID(existingEntry.ID())
			} else {
				k := store.ToKey(treeName)
				if table, err := r.cacheCtx.ranges.Get(ctx, k); err == nil {
					if err := table.Release(uint32(existingEntry.ID().ID())); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func reclaimIDFromExisitingEntries(existingEntries map[string]tree.Entries, id uint32) (*uint32, string) {
	for treeName, existingEntries := range existingEntries {
		for _, existingEntry := range existingEntries {
			if id == uint32(existingEntry.ID().ID()) {
				return &id, treeName
			}
		}
	}
	return nil, ""
}

func claimIDFromExisitingEntries(existingEntries map[string]tree.Entries) (*uint32, string) {
	for treeName, existingEntries := range existingEntries {
		for _, existingEntry := range existingEntries {
			return ptr.To[uint32](uint32(existingEntry.ID().ID())), treeName
		}
	}
	return nil, ""
}

func isReserved(parentName, index string) bool {
	return parentName == fmt.Sprintf("%s.%s", index, vxlanbev1alpha1.VXLANIndexReservedMaxName) ||
		parentName == fmt.Sprintf("%s.%s", index, vxlanbev1alpha1.VXLANIndexReservedMinName)
}
