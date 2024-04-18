default: 
	build

build:
	@docker-compose build

run:
	@docker-compose up -d

down:
	@docker-compose down --remove-orphans
t:
	@go test -cover ./...