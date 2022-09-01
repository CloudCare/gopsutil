//go:build darwin && !cgo
// +build darwin,!cgo

package disk

import (
	"context"

	"github.com/CloudCare/gopsutil/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
