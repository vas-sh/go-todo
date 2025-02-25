lint:
	golangci-lint run -v -c ./.golangcli.yaml ./...
run:
	go run cmd/main.go
integration-test:
	go test ./integration-tests/... -count=1
unit-test:
	go test ./internal/... -count=1
create-db:
	sudo docker run -e POSTGRES_DB=tododb -e POSTGRES_USER=todouser -e POSTGRES_PASSWORD=2222 -d -p 5432:5432 postgres:17.4-alpine3.21
start-db:
	sudo docker start b80def1bbda0
stop-db:
	sudo docker stop b80def1bbda0
install-mock:
	go install go.uber.org/mock/mockgen@latest
gen:
	go generate ./...