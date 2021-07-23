HTTP_NAME 		    := http
MIGRATE_NAME   		:= migrate
PKG            		:= github.com/advita-comics/advita-comics-backend
PKG_LIST       		:= $(shell go list ${PKG}/... | grep -v /vendor/)
NAMESPACE	   		:= "default"

all: setup test build

setup: ## Installing all service dependencies.
	echo "Setup..."
	GO111MODULE=on go mod vendor

.PHONY: config
config: ## Creating the local config yml.
	echo "Creating local config yml ..."
	cp config.example.yml local.yml

db\:migrate: ## Run migrations.
	cd cmd/$(MIGRATE_NAME) && go build && ./$(MIGRATE_NAME) -config=../../local.yml -migrate-path=../../db/migrations

build: ## Build the executable file of service.
	echo "Building..."
	cd cmd/$(HTTP_NAME) && go build

run: build ## Run service with local config.
	echo "Running..."
	cd cmd/$(HTTP_NAME) && ./$(HTTP_NAME) -config=../../local.yml

lint: ## Run lint for all packages.
	echo "Linting..."

	golangci-lint run

test: ## Run tests for all packages.
	echo "Testing..."
	go test -race -count=1 ${PKG_LIST}

coverage: ## Calculating code test coverage.
	echo "Calculating coverage..."
	PKG=$(PKG) ./tools/coverage.sh

clean: ## Cleans the temp files and etc.
	echo "Clean..."
	rm -f cmd/$(HTTP_NAME)/$(HTTP_NAME)

help: ## Display this help screen
	grep -E '^[a-zA-Z_\-\:]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ": .*?## "}; {gsub(/[\\]*/,""); printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: imports
imports:
	goimports -l -w .
