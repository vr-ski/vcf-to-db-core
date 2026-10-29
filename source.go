package core

import "context"

// Source reads generic records from an input source.
//
// Implementations must return io.EOF when no records remain.
// Implementations must honor context cancellation where practical.
type Source interface {
	Next(ctx context.Context) (Record, error)
	Close() error
}
