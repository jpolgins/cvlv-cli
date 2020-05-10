DC ?= docker-compose
DCE ?= docker-compose exec
.DEFAULT_GOAL := help

up: ## [Docker] Build, (re)create, and start application in docker.
	$(DC) up -d && $(DCE) app go mod download

run: ## [Docker] Run the app
	$(DCE) app go run cmd/main.go

ps: ## [Docker] List containers
	$(DC) ps

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

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up run ps resart rm check redis-clear redis-clear help
