package core

// Record represents a single logical data record.
//
// Field names are source-independent identifiers. Their interpretation
// and mapping to target fields are defined outside the Record type.
type Record map[string]Value

// Value represents a generic field value.
//
// The underlying value must be a type supported by the ETL core and its
// adapters. Supported types are defined by the core's value contract.
type Value any
