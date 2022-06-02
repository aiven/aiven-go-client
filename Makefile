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
testacc:
	AIVEN_ACC=1 go test -v -count 1 -parallel 20 -ginkgo.v -ginkgo.progress -ginkgo.focus=$(FOCUS) --cover -timeout 30m .

test:
	go test -v --cover ./... -timeout 15m

lint:
	./bin/golangci-lint run --no-config -E gofmt --issues-exit-code=0 --timeout=30m ./...