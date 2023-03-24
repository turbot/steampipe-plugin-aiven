package aiven

import (
	"context"
	"errors"
	"path"

	aivenClient "github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for Aiven API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		var ar aivenClient.Error
		if errors.As(err, &ar) {
			// Added to support regex in not found errors
			for _, pattern := range notFoundErrors {
				if ok, _ := path.Match(pattern, ar.Error()); ok {
					return true
				}
			}
		}
		return false
	}
}
