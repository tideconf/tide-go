projectname?=tide-go

default: help

.PHONY: help clean test

help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: clean ## display test coverage
	go test -json -v ./... | gotestfmt
