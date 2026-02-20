include .env
export $(shell sed 's/=.*//' .env)

start:
	@go run app.go
lint:
	@golangci-lint run
tests:
	@go test -v ./test/...
tests-%:
	@go test -v ./test/... -run=$(shell echo $* | sed 's/_/./g')
testsum:
	@cd test && gotestsum --format testname
swagger:
	@cd src && swag init
migration-%:
	@migrate create -ext sql -dir migrations create-table-$(subst :,_,$*)
migrate-up:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path ./migrations up
migrate-down:
	@migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path ./migrations down
migrate-docker-up:
	@docker run -v ./migrations:/migrations --network fiber-boilerplate_fiber-boilerplate-network migrate/migrate -path=/migrations/ -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up
	@docker rm -f $$(docker ps -a -q -f ancestor=migrate/migrate) 2>/dev/null || true
migrate-docker-down:
	@docker run -v ./migrations:/migrations --network fiber-boilerplate_fiber-boilerplate-network migrate/migrate -path=/migrations/ -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down -all
	@docker rm -f $$(docker ps -a -q -f ancestor=migrate/migrate) 2>/dev/null || true
seeder-local:
	@go build -o cmd/seeder/main ./seeder
	@TAG=all USER=superuser ./cmd/seeder/main
seeder-docker:
	@docker compose run --rm -e TAG=all -e USER=superuser fiber-boilerplate ./seeder
docker:
	@docker compose up --build
docker-test:
	@docker compose up -d && make tests
docker-down:
	@docker compose down --rmi all --volumes --remove-orphans
docker-cache:
	@docker builder prune -f