.PHONY: build build-prod

build: ## Build docker image
	docker-compose build

build-prod: ## Build docker image (production)
	docker-compose -f docker-compose.prod.yml build
