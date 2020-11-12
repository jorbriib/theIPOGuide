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

## Halts the application by stopping all services
down:
	@docker-compose down --remove-orphans

## Execute migrations to the DB
db_migrate:
	@docker-compose exec api sh /database/migrate.sh

## Seeds test data to the DB
db_seed_test:
	@docker-compose exec api sh /database/clean_database.sh
	@docker-compose exec api sh /database/migrate.sh
	@docker-compose exec api sh /database/seed_test.sh

## Re-starts all application services
restart: down up

## Prunes all volumes and images
prune:
	@docker-compose down -v
	@docker rmi theipoguide_frontend
	@docker rmi theipoguide_api

## Runs tests for src using local code
test:
	@docker-compose -f docker-compose.test.yml up --abort-on-container-exit

## Builds application services images
build:
	@docker-compose build

## Build docker production images with tag
build_prod:
	@docker build --target production -t theipoguide_backend_prod -f backend/Dockerfile backend
	@docker build --target migrations -t theipoguide_migrations_prod -f database/Dockerfile database

## Push docker production images to ECR
push_prod:
	@$$(aws ecr get-login --no-include-email --region ${AWS_REGION})
	@docker tag theipoguide_backend_prod:latest ${AWS_ECR_URL}/${AWS_ECR_REPO}:latest
	@docker push ${AWS_ECR_URL}/${AWS_ECR_REPO}:latest
	@docker tag theipoguide_migrations_prod:latest ${AWS_ECR_URL}/${AWS_ECR_MIGRATIONS_REPO}:latest
	@docker push ${AWS_ECR_URL}/${AWS_ECR_MIGRATIONS_REPO}:latest

deploy_backend_prod:
	sh backend/deploy.sh

deploy_web_prod:
	sh web/deploy.sh