//go:build nsync_baseline

package nsync

import "github.com/apsis-io/nsync/internal/baseline"

var (
	newBenchTryMutex         = baseline.NewTryMutex
	newBenchSemaphore        = baseline.NewSemaphore
	newBenchNamedMutex       = baseline.NewNamedMutex
	newBenchNamedOnceMutex   = baseline.NewNamedOnceMutex
	newBenchControlWaitGroup = baseline.NewControlWaitGroup
)

type benchOnceMutex = baseline.OnceMutex
type benchSyncFlag = baseline.SyncFlag
