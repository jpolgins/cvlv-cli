DC ?= docker-compose
DCE ?= docker-compose exec
.DEFAULT_GOAL := help

up: ## [Docker] Build, (re)create, and start application in docker.
	$(DC) up -d && $(DCE) app go mod download

run: ## [Docker] Run the app
	$(DCE) app go run cmd/*.go

ps: ## [Docker] List containers
	$(DC) ps

ssh: ## [Docker] SSH into app container
	$(DCE) app sh

stop: ## [Docker] Stop containers
	$(DC) stop

restart: ## [Docker] Restart containers
	$(DC) restart

rm: ## [Docker] Remove containers and its volumes
	$(DC) down -v --remove-orphans

redis-cli: ## [Redis] Redis CLI
	$(DCE) redis redis-cli

redis-clear: ## [Redis] Delete all the keys of all the existing databases
	$(DCE) redis redis-cli FLUSHALL ASYNC

fmt: ## [Go] Run go fmt
	$(DCE) app gofmt -s

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up run ps ssh resart rm check redis-clear redis-clear help
