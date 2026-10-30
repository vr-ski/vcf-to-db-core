package core

import (
	"context"
	"io"
)

// SchemaManager provisions the target database using the supplied schema.
//
// Recreate must create a fresh database and schema, replacing any existing
// target database. Implementations are responsible for database-specific
// lifecycle operations.
type SchemaManager interface {
	Recreate(ctx context.Context, schema Schema) error
}

// RecordWriter writes batches of records into a specified target table.
//
// Implementations must honor context cancellation where practical.
// Database-specific transaction and conflict behavior belongs to the adapter.
type RecordWriter interface {
	WriteBatch(ctx context.Context, table string, records []Record) error
}

// Target represents a database target.
//
// It combines schema provisioning and record writing while keeping
// those responsibilities independently implementable and testable.
type Target interface {
	SchemaManager
	RecordWriter
	io.Closer
}
