lint:
	golangci-lint run -v -c ./.golangcli.yaml ./...
run:
	go run cmd/main.go
integration-test:
	go test ./integration-tests/... -count=1
unit-test:
	go test ./internal/... -count=1
