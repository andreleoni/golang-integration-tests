mocks:
	docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage
	chown -R $USER:$USER *

unit_test:
	go test ./...

integration_test:
	go test -tags=integration ./...
