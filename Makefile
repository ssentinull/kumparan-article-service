run-server:
	@modd -f ./.modd/server.modd.conf

run-migrate-up:
	@go run cmd/migration/main.go -mode=up

run-migrate-down:
	@go run cmd/migration/main.go -mode=down

run-test:
	@go test -v -cover -covermode=atomic ./pkg/...