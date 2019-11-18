image: ## Build docker image
	@ docker-compose build

start: ## Start docker containers
	@ docker-compose up -d

status: ## Status of containers
	@ docker-compose ps

stop: ## Stop containers
	@ docker-compose stop

clean:stop ## Stop containers and clean data and workspace
	@ docker-compose down -v --remove-orphans

test: ## Run tests
	@ docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	@ docker-compose -f docker-compose.test.yml down --volumes
