GOLANGCI_LINT_VERSION = v1.33.0
GO = go
TOOLSDIR = $(TMPDIR)/tos-aoc
GOLANGCI_LINT = ${TOOLSDIR}/golangci-lint

all: test ## just executes test

test: lint ## executes tests and output the puzzle results
	${GO} test -v ./... | grep "AdventOfCode"

lint: ## Executes all linters
	${GOLANGCI_LINT} run --enable-all --exclude-use-default=false --disable=paralleltest

help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


setupLinter:
	mkdir -p ${TOOLSDIR}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLSDIR) $(GOLANGCI_LINT_VERSION)



.DEFAULT_GOAL := all