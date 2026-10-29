// Package core defines the database-independent contracts and domain types
// used by the VCF-to-database ETL application.
//
// The package contains no database drivers, file-format parsers, or
// infrastructure-specific implementations. Source and target adapters
// implement the contracts defined here.
//
// The application uses these contracts to read generic records, map their
// values to a user-defined target schema, and write them to a database.
//
// Implementations of external-system interfaces belong in separate
// adapter modules.
package core
