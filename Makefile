#!make

.DEFAULT_GOAL := help
.PHONY: help

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=35

## Shows help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Starts all app services to run the application (hot-reload enabled)
up:
	@docker-compose up -d && docker-compose logs -f

## Halts the application by stoping all services
down:
	@docker-compose down --remove-orphans

## Re-starts all application services
restart: down up

## Builds application services images for e2e suite
build:
	@docker-compose build

## Runs tests for backend using local code and tag $IMAGE_TAG
test:
	@docker-compose -f docker-compose.test.yml up --abort-on-container-exit
