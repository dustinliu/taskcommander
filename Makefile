app := taskcommander
cli := tc

export TC_RUNTIME_ENV=dev

generate:
	@echo 'generate mock codes'
	@go generate ./...

debug-build: generate
	@echo building debug version...
	@go build -gcflags="all=-N -l" -ldflags "-X main.version=`cat version`" -o $(app) ./cmd/gui/$(app).go
	@go build -gcflags="all=-N -l" -ldflags "-X main.version=`cat version`" -o $(cli) ./cmd/cli/$(cli).go

run: generate
	@go run cmd/taskcommander.go

lint:
	@golangci-lint run

test: generate
	@echo testing...
	@go test -timeout 10s ./...
	@echo

clean:
	@go clean
	@go clean -testcache
	@rm -f $(app) $(cli)

release:
	git tag `cat version`
	git push origin `cat version`

.PHONY: debug-build clean test lint generate
