.PHONY: test run fmt vet

test:
	go test ./...

run:
	go run ./cmd/migrator -config config/runtime.yaml

fmt:
	gofmt -w cmd internal

vet:
	go vet ./...
