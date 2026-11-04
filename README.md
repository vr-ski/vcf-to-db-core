# VCF-to-DB Core

Database-independent core library for the VCF-to-database ETL application.

## Purpose

`vcf-to-db-core` defines the domain models, contracts, and application-level functionality shared by source and target adapters.

It contains no database drivers, database-specific SQL, or source-format-specific parsing implementations.

## Architecture

The project follows Clean Architecture principles and uses a ports-and-adapters approach.

```text
                 +----------------------+
                 |      ETL Core        |
                 |                      |
                 | - Domain models      |
                 | - Mapping contracts  |
                 | - Validation         |
                 | - Adapter interfaces|
                 +----------+-----------+
                            |
              +-------------+-------------+
              |                           |
     +--------v---------+       +---------v----------+
     |   Source Adapter |       |   Target Adapter   |
     |                  |       |                    |
     | VCF Parser       |       | PostgreSQL         |
     +------------------+       +--------------------+
```

The core owns the interfaces. Adapters implement them and depend on the core module.

## Repository Responsibilities

### Included

- Generic records and values.
- Logical database schema definitions.
- Mapping definitions.
- Source and target interfaces.
- Adapter capability declarations.
- Database-independent validation.
- Common errors and contracts.

### Excluded

- VCF parsing and normalization.
- Database drivers and SQL generation.
- Database creation and lifecycle implementation.
- Configuration file parsing.
- Infrastructure-specific logging.
- CLI and deployment orchestration.

## Current API

| File | Responsibility |
|---|---|
| `record.go` | Generic records and values |
| `schema.go` | Logical schema definitions |
| `mapping.go` | Source-to-target mapping definitions |
| `source.go` | Source adapter interface |
| `target.go` | Target adapter interfaces |
| `capabilities.go` | Adapter capabilities |
| `errors.go` | Shared error definitions |
| `validation.go` | Schema and mapping validation |

## Adapter Repositories

The initial adapters are maintained separately:

- `vcf-to-db-postgresql`: PostgreSQL target adapter.
- `vcf-to-db-source-vcf`: VCF source adapter (repository name to be finalized).

Adapters should depend on a released version of this module.

## Development

Run tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Format source files:

```bash
gofmt -w .
```

## Design Principles

1. Keep the core independent of infrastructure.
2. Keep interfaces small and consumer-oriented.
3. Prefer explicit contracts over implicit behavior.
4. Validate configuration and schemas before destructive operations.
5. Avoid abstractions without a concrete requirement.
6. Keep database-specific behavior inside target adapters.
7. Preserve deterministic and testable application behavior.

## Project Status

Early development — public contracts are being established.

The API is not yet stable and may change before the first release.
