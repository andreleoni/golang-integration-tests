generate_mocks:
	docker run -v "$PWD":/src -w /src vektra/mockery --all

unit_test:
	go test ./...

integration_test:
	go test -tags=integration ./...
