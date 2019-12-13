ci: lint
.PHONY: ci

#################################################
# Bootstrapping for base golang package deps
#################################################
bootstrap:
	go mod download
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0
	./bin/golangci-lint --version

#################################################
# Testing and linting
#################################################
test:
	go test -v ./...

lint:
	go fmt  ./...
	./bin/golangci-lint run --no-config --issues-exit-code=0 ./...