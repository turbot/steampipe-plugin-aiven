package aiven

import (
	"context"
	"path"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for Aiven API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range notFoundErrors {
			// handle not found error
			if ok, _ := path.Match(pattern, "404"); ok {
				return true
			}
		}
		return false
	}
}
