# Always turn on go module when use `go` command.
GO := GO111MODULE=on go

# GO build preparation
.PHONY: ci
ci:
	$(GO) mod download && \
	$(GO) mod verify && \
	$(GO) mod vendor && \
	$(GO) fmt ./... \

# Build GO application
# -mod=vendor 
# tells the go command to use the vendor directory. In this mode,
# the go command will not use the network or the module cache.
# -v
# print the names of packages as they are compiled.
# -a
# force rebuilding of packages that are already up-to-date.
# -o
# -ldsflags
# tells the version and go version.
.PHONY: build
build:
	$(GO) build -mod=vendor -ldflags '$(LDFLAGS)' -a -v -o $(GO_BINARY_NAME) ./cmd/main.go

start:
	go run ./cmd/main.go serve-rest

swag:
	swag init -g ./cmd/cmds/rest.go -o ./docs

lint:
	go fmt ./...
	if  [ $(shell which golangcli-lint)=="golangcli-lint not found" ]; then \
		echo "golangcli-lint not found!"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2; \
	fi
		echo "golangcli-lint found!"; \
		golangci-lint run

.PHONY: test
test:
	go test ./... -coverprofile coverage.out

# Clean up when build the application on local directory.
.PHONY: clean
clean:
	@rm -rf $(GO_BINARY_NAME) ./vendor