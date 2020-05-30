DC ?= docker-compose
DCE ?= docker-compose exec
.DEFAULT_GOAL := help

run: ## [Docker] Run the app
	$(DC) up -d && $(DCE) app go mod download && $(DCE) app go run cmd/*.go

ps: ## [Docker] List containers
	$(DC) ps

ssh: ## [Docker] SSH into app container
	$(DCE) app sh

clean: ## [Docker] Remove containers and its volumes
	$(DC) down -v --remove-orphans

redis-cli: ## [Redis] Redis CLI
	$(DCE) redis redis-cli

redis-flush: ## [Redis] Delete all the keys of all the existing databases
	$(DCE) redis redis-cli FLUSHALL ASYNC

test: ## [Go] Run tests
	$(DCE) app go test ./...

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: ssh test help
