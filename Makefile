image: ## Build docker image
	@ docker-compose build

start: ## Start docker containers
	@ docker-compose up -d

stop: ## Stop containers
	@ docker-compose stop

clean:stop ## Stop containers and clean data and workspace
	@ docker-compose down -v --remove-orphans
