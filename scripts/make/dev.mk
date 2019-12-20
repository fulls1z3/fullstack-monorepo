.PHONY: status start stop clean

status: ## Get status of containers
	docker-compose ps

start: ## Start docker containers
	docker-compose up -d

stop: ## Stop docker containers
	docker-compose stop

clean:stop ## Stop docker containers, clean data and workspace
	docker-compose down -v --remove-orphans
